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
    "clamor/models"
)

// Get Campaign results. List of campaigns
type GetCampaiagnsResult struct {
    Campaigns []models.Campaign
}

// Get Campaign Types. List of campaign types.
type GetCampaiagnTypesResult struct {
    CampaignTypes []models.CampaignType
}

// Get About response. Just version strings and similar things. (It may make sense to use a standard format for all services?)
type GetAboutResponse struct {
    Name string
    CoreVersion string
    Version string
    BuildTime string
}
