// Copyright 2018~now John Deng (hi.devops.io@gmail.com).
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
	"github.com/go-openapi/spec"
	"hidevops.io/hiboot/pkg/at"
	"hidevops.io/hiboot/pkg/system"
)

type ApiInfoBuilderInterface interface {
	Title(value string) ApiInfoBuilderInterface
	Description(value string) ApiInfoBuilderInterface
	TermsOfServiceUrl(value string) ApiInfoBuilderInterface
	Version(value string) ApiInfoBuilderInterface
	ContactName(value string) ApiInfoBuilderInterface
	ContactEmail(value string) ApiInfoBuilderInterface
	ContactURL(value string) ApiInfoBuilderInterface
	Contact(value Contact) ApiInfoBuilderInterface
	LicenseName(value string) ApiInfoBuilderInterface
	LicenseURL(value string) ApiInfoBuilderInterface
	License(value License) ApiInfoBuilderInterface
	Schemes(values ...string) ApiInfoBuilderInterface
	Host(values string) ApiInfoBuilderInterface
	BasePath(values string) ApiInfoBuilderInterface
}

type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type License struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type ApiInfo struct {
	at.ConfigurationProperties `value:"swagger"`
	spec.Swagger

	SystemServer *system.Server
	AppVersion   string `value:"${app.version}"`
}

func ApiInfoBuilder() ApiInfoBuilderInterface {
	return &ApiInfo{
		Swagger: spec.Swagger{
			SwaggerProps: spec.SwaggerProps{
				Swagger: "2.0",
				Info:    &spec.Info{},
				Paths: &spec.Paths{
					VendorExtensible: spec.VendorExtensible{},
					Paths:            make(map[string]spec.PathItem),
				},
			},
		},
	}
}

func (b *ApiInfo) Title(value string) ApiInfoBuilderInterface {
	b.Swagger.SwaggerProps.Info.Title = value
	return b
}

func (b *ApiInfo) Description(value string) ApiInfoBuilderInterface {
	b.Swagger.SwaggerProps.Info.Description = value
	return b
}

func (b *ApiInfo) TermsOfServiceUrl(value string) ApiInfoBuilderInterface {
	b.Swagger.SwaggerProps.Info.TermsOfService = value
	return b
}

func (b *ApiInfo) Version(value string) ApiInfoBuilderInterface {
	b.Swagger.SwaggerProps.Info.Version = value
	return b
}

func (b *ApiInfo) Host(values string) ApiInfoBuilderInterface {
	b.Swagger.SwaggerProps.Host = values
	return b
}

func (b *ApiInfo) BasePath(values string) ApiInfoBuilderInterface {
	b.Swagger.SwaggerProps.BasePath = values
	return b
}

func (b *ApiInfo) Schemes(values ...string) ApiInfoBuilderInterface {
	b.Swagger.SwaggerProps.Schemes = values
	return b
}


func (b *ApiInfo) ContactName(value string) ApiInfoBuilderInterface {
	b.ensureContact()
	b.Swagger.SwaggerProps.Info.Contact.Name = value
	return b
}

func (b *ApiInfo) ensureContact() {
	if b.Swagger.SwaggerProps.Info.Contact == nil {
		b.Swagger.SwaggerProps.Info.Contact = &spec.ContactInfo{}
	}
}

func (b *ApiInfo) ContactEmail(value string) ApiInfoBuilderInterface {
	b.ensureContact()
	b.Swagger.SwaggerProps.Info.Contact.Email = value
	return b
}

func (b *ApiInfo) ContactURL(value string) ApiInfoBuilderInterface {
	b.ensureContact()
	b.Swagger.SwaggerProps.Info.Contact.URL = value
	return b
}

func (b *ApiInfo) Contact(value Contact) ApiInfoBuilderInterface {
	b.ContactName(value.Name)
	b.ContactEmail(value.Email)
	b.ContactURL(value.URL)
	return b
}

func (b *ApiInfo) ensureLicense() {
	if b.Swagger.SwaggerProps.Info.License == nil {
		b.Swagger.SwaggerProps.Info.License = &spec.License{}
	}
}

func (b *ApiInfo) LicenseName(value string) ApiInfoBuilderInterface {
	b.ensureLicense()
	b.Swagger.SwaggerProps.Info.License.Name = value
	return b
}

func (b *ApiInfo) LicenseURL(value string) ApiInfoBuilderInterface {
	b.ensureLicense()
	b.Swagger.SwaggerProps.Info.License.URL = value
	return b
}

func (b *ApiInfo) License(value License) ApiInfoBuilderInterface {
	b.LicenseName(value.Name)
	b.LicenseURL(value.URL)
	return b
}
