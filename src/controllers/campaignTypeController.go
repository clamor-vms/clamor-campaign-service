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

    clamor "github.com/clamor-vms/clamor-go-core"

    "clamor/services"
)

type CampaignTypeController struct {
    campaignService services.ICampaignService
}
func NewCampaignTypeController(campaignService services.ICampaignService) *CampaignTypeController {
    return &CampaignTypeController{
        campaignService: campaignService,
    }
}
func (p *CampaignTypeController) Get(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    campaignTypes, err := p.campaignService.GetCampaignTypes()

    if err == nil {
        //return list of all campaign types
        return clamor.ControllerResponse{Status: http.StatusOK, Body: GetCampaiagnTypesResult{CampaignTypes: campaignTypes}}
    } else {
        //Log the error and send a 500 back.
        log.Output(1, fmt.Sprintf("Campaign Type Controller Get Error: %s", err.Error()))
        return clamor.ControllerResponse{Status: http.StatusInternalServerError, Body: clamor.EmptyResponse{}}
    }
}
//return 404 for all other actions. This is read only static content.
func (p *CampaignTypeController) Post(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    return clamor.ControllerResponse{Status: http.StatusNotFound, Body: clamor.EmptyResponse{}}
}
func (p *CampaignTypeController) Put(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    return clamor.ControllerResponse{Status: http.StatusNotFound, Body: clamor.EmptyResponse{}}
}
func (p *CampaignTypeController) Delete(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    return clamor.ControllerResponse{Status: http.StatusNotFound, Body: clamor.EmptyResponse{}}
}
