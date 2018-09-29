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

package services

import (
    "github.com/jinzhu/gorm"

    "skaioskit/models"
)

type ICampaignService interface {
    CreateCampaign(models.Campaign) models.Campaign
    UpdateCampaign(models.Campaign) models.Campaign
    GetCampaign(string) (models.Campaign, error)
    EnsureCampaignTable()
}

type CampaignService struct {
    db *gorm.DB
}
func NewCampaignService(db *gorm.DB) *CampaignService {
    return &CampaignService{db: db}
}
func (p *CampaignService) CreateCampaign(campaign models.Campaign) models.Campaign {
    p.db.Create(&campaign)
    return campaign
}
func (p *CampaignService) UpdateCampaign(campaign models.Campaign) models.Campaign {
    p.db.Save(&campaign)
    return campaign
}
func (p *CampaignService) GetCampaign(name string) (models.Campaign, error) {
    var campaign models.Campaign
    err := p.db.Where(&models.Campaign{Name: name}).First(&campaign).Error
    return campaign, err
}
func (p *CampaignService) EnsureCampaignTable() {
    p.db.AutoMigrate(&models.Campaign{})
}
func (p *CampaignService) EnsureCampaign(campaign models.Campaign) {
    existing, err := p.GetCampaign(campaign.Name)
    if err != nil {
        p.CreateCampaign(campaign)
    } else {
        existing.Name = campaign.Name
        p.UpdateCampaign(existing)
    }
}
