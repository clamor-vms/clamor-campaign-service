/*
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of the
    License, or (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package controllers

import (
    "net/http"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/services"
)

type CampaignController struct {
    campaignService services.ICampaignService
}
func NewCampaignController(campaignService services.ICampaignService) *CampaignController {
    return &CampaignController{
        campaignService: campaignService,
    }
}
func (p *CampaignController) Get(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *CampaignController) Post(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *CampaignController) Put(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *CampaignController) Delete(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
