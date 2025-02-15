// Copyright 2018 ~ now John Deng (hi.devops.io@gmail.com).
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swagger

import (
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
	"hidevops.io/hiboot/pkg/app/web/context"
	"hidevops.io/hiboot/pkg/at"
	"net/http"
	"path"
)

type Swagger struct {
	at.RestController
	at.RequestMapping `value:"/"`

	apiInfo *ApiInfo
}

func newController(openAPIDefinition *ApiInfo) *Swagger {
	return &Swagger{apiInfo: openAPIDefinition}
}

// TODO: add description 'Implemented by HiBoot Framework'
func (c *Swagger) loadDoc() (retVal []byte, err error) {
	retVal, err = json.MarshalIndent(c.apiInfo.Swagger, "", "  ")
	return
}

func (c *Swagger) serve(ctx context.Context, docsPath string) {
	b, err := c.loadDoc()
	if err == nil {
		// read host dynamically
		c.apiInfo.Swagger.Host = ctx.Host()
		// concat path
		basePath := path.Join(c.apiInfo.Swagger.BasePath, c.RequestMapping.AtValue)

		// get handler
		handler := middleware.Redoc(middleware.RedocOpts{
			BasePath: basePath,
			SpecURL:  path.Join(basePath, "solar-mesh/swagger.json"),
			Path:     docsPath,
		}, http.NotFoundHandler())

		// handle cors
		handler = handlers.CORS()(middleware.Spec(basePath, b, handler))

		// wrap handler
		ctx.WrapHandler(handler)
	}
}

// UI serve static resource via context StaticResource method
func (c *Swagger) Swagger(at struct {
	at.GetMapping `value:"/swagger.json"`
}) (response string) {
	b, err := c.loadDoc()
	if err == nil {
		response = string(b)
	}
	return
}

// UI serve static resource via context StaticResource method
func (c *Swagger) SwaggerUI(at struct {
	at.GetMapping `value:"/swagger-ui"`
}, ctx context.Context) {
	c.serve(ctx, at.GetMapping.AtValue)
	return
}
