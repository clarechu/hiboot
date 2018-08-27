// Copyright 2018 John Deng (hi.devops.io@gmail.com).
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

package jwt

import (
	"github.com/hidevopsio/hiboot/pkg/app/web"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

type JwtController interface {
}

// JwtController is the base web controller that enabled JWT
type Controller struct {
	JwtController
	web.Controller
}

// ParseToken is an util that parsing JWT token from jwt.MapClaims
func (c *Controller) ParseToken(claims jwt.MapClaims, prop string) string {
	return fmt.Sprintf("%v", claims[prop])
}

// ParseToken is an util that parsing JWT token from jwt.MapClaims
func (c *Controller) GetJwtToken() *jwt.Token {
	return c.Ctx.Values().Get("jwt").(*jwt.Token)
}

// GetJwtProperty is an util that parsing JWT token and return single property from jwt.MapClaims
func (c *Controller) GetJwtProperty(propName string) (propVal string) {
	claims := c.GetJwtProperties()
	propVal = fmt.Sprintf("%v", claims[propName])
	return
}

// GetJwtProperty is an util that parsing JWT token and return all properties from jwt.MapClaims
func (c *Controller) GetJwtProperties() (propMap map[string]interface{}) {
	propMap = make(map[string]interface{})

	token := c.GetJwtToken()

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		propMap = claims
	}
	return
}

// GetJwtProperty is an util that parsing JWT token and return all properties in string from jwt.MapClaims
func (c *Controller) GetJwtPropertiesString() (propMap map[string]string) {

	claims := c.GetJwtProperties()

	for name, value := range claims {
		propMap[name] = fmt.Sprintf("%v", value)
	}
	return
}

