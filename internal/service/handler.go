package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"io/ioutil"
	"net/http"

	"github.com/gocraft/web"
	"github.com/icpac-igad/legend-image-generator/internal/legend"
)

type Context struct {
}

func Error(rw web.ResponseWriter, req *web.Request, err interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered panic:", err)
			return
		}
		fmt.Println("no panic recovered")
	}()
}

func (c *Context) HandleGenerateLegend(rw web.ResponseWriter, req *web.Request) {

	// ready request body
	body, err := ioutil.ReadAll(req.Body)

	defer req.Body.Close()

	if err != nil {
		err := appError{Status: http.StatusBadRequest, Message: err.Error()}
		JSONHandleError(rw, err)
		return
	}

	var legendConfig legend.LegendConfig
	err = json.Unmarshal(body, &legendConfig)

	if err != nil {
		JSONHandleError(rw, appError{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	legendImg, err := legend.GetLegendImg(legendConfig)

	if err != nil {
		err := appError{Status: http.StatusBadRequest, Message: err.Error()}
		JSONHandleError(rw, err)
		return
	}

	buff := new(bytes.Buffer)
	err = png.Encode(buff, legendImg)

	if err != nil {
		err := appError{Status: http.StatusBadRequest, Message: "error encoding legend image"}
		JSONHandleError(rw, err)
		return
	}

	rw.Header().Set("Content-Type", "image/png")
	rw.Write(buff.Bytes())
}

func initRouter(basePath string) *web.Router {
	// create router
	router := web.New(Context{})

	// ovveride gocraft defualt error handler
	router.Error(Error)

	// add middlewares
	router.Middleware(loggerMiddleware)
	// router.Middleware(web.ShowErrorsMiddleware)

	// handle routes
	router.Post("/", (*Context).HandleGenerateLegend)

	return router
}
