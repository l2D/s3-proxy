package responsehandler

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/authx/models"
	"github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/config"
	"github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/utils"
	"github.com/pkg/errors"
)

type handler struct {
	req        *http.Request
	res        http.ResponseWriter
	cfgManager config.Manager
	targetKey  string
}

func (h *handler) UpdateRequestAndResponse(req *http.Request, res http.ResponseWriter) {
	// Update request
	h.req = req
	// Update response
	h.res = res
}

func (h *handler) PreconditionFailed() {
	h.res.WriteHeader(http.StatusPreconditionFailed)
}

func (h *handler) NotModified() {
	h.res.WriteHeader(http.StatusNotModified)
}

func (h *handler) NoContent() {
	h.res.WriteHeader(http.StatusNoContent)
}

func (h *handler) TargetList() {
	// Get configuration
	cfg := h.cfgManager.GetConfig()

	// Create data structure
	data := targetListData{
		Request: h.req,
		User:    models.GetAuthenticatedUserFromContext(h.req.Context()),
		Targets: cfg.Targets,
	}

	h.handleGenericAnswer(
		nil,
		data,
		nil,
		nil,
		cfg.Templates.TargetList,
		cfg.Templates.Helpers,
	)
}

func (h *handler) RedirectWithTrailingSlash() {
	//  Get path
	p := h.req.URL.RequestURI()
	// Check if path doesn't start with /
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	// Check if path doesn't end with /
	if !strings.HasSuffix(p, "/") {
		p += "/"
	}
	// Redirect
	http.Redirect(h.res, h.req, p, http.StatusFound)
}

func (h *handler) StreamFile(
	loadFileContent func(ctx context.Context, path string) (string, error),
	input *StreamInput,
) error {
	// Get configuration
	cfg := h.cfgManager.GetConfig()
	// Get target configuration (exists in this case)
	targetCfg := cfg.Targets[h.targetKey]

	// Set headers from object
	setHeadersFromObjectOutput(h.res, input)

	// Check if headers templates are defined in the GET configuration
	if targetCfg.Actions != nil &&
		targetCfg.Actions.GET != nil &&
		targetCfg.Actions.GET.Config != nil &&
		targetCfg.Actions.GET.Config.StreamedFileHeaders != nil {
		// Target template helpers
		var tplHelpers []*config.TargetHelperConfigItem
		// Check if target templates are defined
		if targetCfg.Templates != nil {
			tplHelpers = targetCfg.Templates.Helpers
		}

		// Get template content
		helpersTpl, err := h.loadAllHelpersContent(
			loadFileContent,
			tplHelpers,
			cfg.Templates.Helpers,
		)
		// Check error
		if err != nil {
			return err
		}

		// Create data structure
		data := &streamFileHeaderData{
			Request:    h.req,
			User:       models.GetAuthenticatedUserFromContext(h.req.Context()),
			StreamFile: input,
		}
		// Manage headers
		headers, err := h.manageHeaders(helpersTpl, targetCfg.Actions.GET.Config.StreamedFileHeaders, data)
		// Check error
		if err != nil {
			return err
		}

		// Loop over them to add them to response
		for k, v := range headers {
			// Add them
			h.res.Header().Set(k, v)
		}
	}

	// Copy data stream to output stream
	_, err := io.Copy(h.res, input.Body)

	return errors.WithStack(err)
}

func (h *handler) FoldersFilesList(
	loadFileContent func(ctx context.Context, path string) (string, error),
	entries []*Entry,
) {
	// Get config
	cfg := h.cfgManager.GetConfig()

	// Get target configuration
	targetCfg := cfg.Targets[h.targetKey]

	// Helpers list
	var helpersCfgList []*config.TargetHelperConfigItem

	// Target template config item
	var tplConfigItem *config.TargetTemplateConfigItem

	// Get helpers template configs
	if targetCfg != nil && targetCfg.Templates != nil {
		// Save
		helpersCfgList = targetCfg.Templates.Helpers
		tplConfigItem = targetCfg.Templates.FolderList
	}

	// Create bucket list data for templating
	data := &folderListingData{
		Request:    h.req,
		User:       models.GetAuthenticatedUserFromContext(h.req.Context()),
		Entries:    entries,
		BucketName: targetCfg.Bucket.Name,
		Name:       targetCfg.Name,
	}

	h.handleGenericAnswer(
		loadFileContent,
		data,
		tplConfigItem,
		helpersCfgList,
		cfg.Templates.FolderList,
		cfg.Templates.Helpers,
	)
}

func (h *handler) handleGenericAnswer(
	loadFileContent func(ctx context.Context, path string) (string, error),
	data interface{},
	tplCfgItem *config.TargetTemplateConfigItem,
	helpersTplCfgItems []*config.TargetHelperConfigItem,
	baseTpl *config.TemplateConfigItem,
	helpersTplFilePathList []string,
) {
	// Get helpers template content
	helpersContent, err := h.loadAllHelpersContent(
		loadFileContent,
		helpersTplCfgItems,
		helpersTplFilePathList,
	)
	// Check if error exists
	if err != nil {
		// Return an internal server error
		h.InternalServerError(
			loadFileContent,
			err,
		)

		return
	}

	// Save in template
	tplContent := helpersContent

	// Check if a target template configuration exists
	// Note: Done like this and not with list to avoid creating list of 1 element
	// and to avoid loops etc to save potential memory and cpu
	if tplCfgItem != nil {
		// Load template content
		tpl, err2 := h.loadTemplateContent(
			loadFileContent,
			tplCfgItem,
		)
		// Concat
		tplContent = tplContent + "\n" + tpl
		// Save error
		err = err2
	} else {
		// Get template from general configuration
		tpl, err2 := loadLocalFileContent(baseTpl.Path)
		// Concat
		tplContent = tplContent + "\n" + tpl
		// Save error
		err = err2
	}

	// Check if error exists
	if err != nil {
		// Return an internal server error
		h.InternalServerError(
			loadFileContent,
			err,
		)

		return
	}

	// Store headers
	var headers map[string]string

	// Create header data
	hData := &genericHeaderData{
		Request: h.req,
		User:    models.GetAuthenticatedUserFromContext(h.req.Context()),
	}

	// Check if target config item exists
	if tplCfgItem != nil {
		// Manage headers
		headers, err = h.manageHeaders(
			helpersContent,
			tplCfgItem.Headers,
			hData,
		)
	} else {
		// Manage headers
		headers, err = h.manageHeaders(
			helpersContent,
			baseTpl.Headers,
			hData,
		)
	}

	// Check if error exists
	if err != nil {
		// Return an internal server error
		h.InternalServerError(
			loadFileContent,
			err,
		)

		return
	}

	// Execute main template
	bodyBuf, err := utils.ExecuteTemplate(tplContent, data)
	// Check error
	if err != nil {
		h.InternalServerError(loadFileContent, err)

		return
	}

	// Manage status code
	statusCode, err := h.manageStatus(helpersContent, tplCfgItem, baseTpl.Status, data)
	// Check error
	if err != nil {
		h.InternalServerError(loadFileContent, err)

		return
	}

	// Send
	err = h.send(bodyBuf, headers, statusCode)
	// Check error
	if err != nil {
		// Return an internal server error
		h.InternalServerError(
			loadFileContent,
			err,
		)
	}
}
