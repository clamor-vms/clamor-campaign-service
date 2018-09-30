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
    "log"
    "strconv"
    "net/http"
    "encoding/json"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/models"
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
    idStr, ok := r.URL.Query()["id"]

    if ok {
        id, err := strconv.ParseUint(idStr[0], 10, 32)

        if err == nil {
            campaign, err := p.campaignService.GetCampaign(uint(id))

            if err == nil {
                return skaioskit.ControllerResponse{Status: http.StatusOK, Body: campaign}
            }
        }
    } else {
        campaigns, err := p.campaignService.GetCampaigns()

        if err == nil {
            return skaioskit.ControllerResponse{Status: http.StatusOK, Body: GetCampaiagnsResult{Campaigns: campaigns}}
        }
    }

    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *CampaignController) Post(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    log.Output(1, "here")

    //Parse request into struct
    decoder := json.NewDecoder(r.Body)
    var data models.Campaign
    err := decoder.Decode(&data)

    if err == nil {
        campaign, err := p.campaignService.CreateCampaign(data)
        if err == nil {
            return skaioskit.ControllerResponse{Status: http.StatusOK, Body: campaign}
        } else {
            return skaioskit.ControllerResponse{Status: http.StatusInternalServerError, Body: skaioskit.EmptyResponse{}}
        }
    } else {
        //if json doesn't map to struct return error
        return skaioskit.ControllerResponse{Status: http.StatusInternalServerError, Body: skaioskit.EmptyResponse{}}
    }
}
func (p *CampaignController) Put(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    //Parse request into struct
    decoder := json.NewDecoder(r.Body)
    var data models.Campaign
    err := decoder.Decode(&data)

    if err != nil {
        //if json doesn't map to struct return error
        return skaioskit.ControllerResponse{Status: http.StatusInternalServerError, Body: skaioskit.EmptyResponse{}}
    } else {
        campaign, err := p.campaignService.UpdateCampaign(data)
        if err == nil {
            return skaioskit.ControllerResponse{Status: http.StatusOK, Body: campaign}
        } else {
            return skaioskit.ControllerResponse{Status: http.StatusInternalServerError, Body: skaioskit.EmptyResponse{}}
        }
    }
}
func (p *CampaignController) Delete(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}