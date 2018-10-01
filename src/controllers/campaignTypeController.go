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
    "fmt"
    "log"
    "net/http"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/services"
)

type CampaignTypeController struct {
    campaignService services.ICampaignService
}
func NewCampaignTypeController(campaignService services.ICampaignService) *CampaignTypeController {
    return &CampaignTypeController{
        campaignService: campaignService,
    }
}
func (p *CampaignTypeController) Get(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    campaignTypes, err := p.campaignService.GetCampaignTypes()

    if err == nil {
        return skaioskit.ControllerResponse{Status: http.StatusOK, Body: GetCampaiagnTypesResult{CampaignTypes: campaignTypes}}
    } else {
        log.Output(1, fmt.Sprintf("Campaign Type Controller Get Error: %s", err.Error()))
        return skaioskit.ControllerResponse{Status: http.StatusInternalServerError, Body: skaioskit.EmptyResponse{}}
    }
}
func (p *CampaignTypeController) Post(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *CampaignTypeController) Put(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *CampaignTypeController) Delete(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
