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

    "clamor/models"
)

type ICampaignService interface {
    CreateCampaign(models.Campaign) (models.Campaign, error)
    UpdateCampaign(models.Campaign) (models.Campaign, error)
    GetCampaign(uint) (models.Campaign, error)
    GetCampaigns() ([]models.Campaign, error)
    DeleteCampaign(uint) (error)
    EnsureCampaignTable()
    GetCampaignTypes() ([]models.CampaignType, error)
    EnsureCampaignTypeTable()
    EnsureCampaignTypes()
}

type CampaignService struct {
    db *gorm.DB
}
func NewCampaignService(db *gorm.DB) *CampaignService {
    return &CampaignService{db: db}
}
//campaigns
func (p *CampaignService) CreateCampaign(campaign models.Campaign) (models.Campaign, error) {
    err := p.db.Create(&campaign).Error
    
    return campaign, err
}
func (p *CampaignService) UpdateCampaign(campaign models.Campaign) (models.Campaign, error) {
    err := p.db.Save(&campaign).Error
    
    return campaign, err
}
func (p *CampaignService) GetCampaign(id uint) (models.Campaign, error) {
    var campaign models.Campaign
    
    err := p.db.Preload("CampaignType").First(&campaign, id).Error

    return campaign, err
}
func (p *CampaignService) GetCampaigns() ([]models.Campaign, error) {
    var campaigns []models.Campaign
    
    err := p.db.Preload("CampaignType").Find(&campaigns).Error

    return campaigns, err
}
func (p *CampaignService) DeleteCampaign(id uint) error {
    campaign, err := p.GetCampaign(id)

    if err == nil {
        err = p.db.Delete(&campaign).Error
    }

    return err
}
func (p *CampaignService) EnsureCampaignTable() {
    p.db.AutoMigrate(&models.Campaign{})
    p.db.Model(&models.Campaign{}).AddForeignKey("campaign_type_id", "campaign_type(id)", "CASCADE", "RESTRICT")
}
//campaignTypes
func (p *CampaignService) GetCampaignTypes() ([]models.CampaignType, error) {
    var campaignTypes []models.CampaignType
    
    err := p.db.Find(&campaignTypes).Error
    
    return campaignTypes, err
}
func (p *CampaignService) EnsureCampaignTypeTable() {
    p.db.AutoMigrate(&models.CampaignType{})
}
func (p *CampaignService) EnsureCampaignTypes() {
    p.ensureCampaignType("US House Congress Race", "A campaign to elect a federal canididate to the US House.")
    p.ensureCampaignType("US Senate Congress Race", "A campaign to elect a federal canididate to the US Senate.")
    p.ensureCampaignType("Ballot Initiative", "A campaign to collection ballot initiative signature of register voters.")
}
func (p *CampaignService) ensureCampaignType(name string, description string) {
    existing, err := p.getCampaignType(name)
    if err == nil {
        existing.Description = description;
        p.updateCampaignType(existing)
    } else {
        data := models.CampaignType{Name: name, Description: description}
        p.createCampaignType(data)
    }
}
func (p *CampaignService) createCampaignType(campaignType models.CampaignType) (models.CampaignType, error) {
    err := p.db.Create(&campaignType).Error
    
    return campaignType, err
}
func (p *CampaignService) updateCampaignType(campaignType models.CampaignType) (models.CampaignType, error) {
    err := p.db.Save(&campaignType).Error
    
    return campaignType, err
}
func (p *CampaignService) getCampaignType(name string) (models.CampaignType, error) {
    var campaignType models.CampaignType
    
    err := p.db.Where(&models.CampaignType{Name: name}).First(&campaignType).Error
    
    return campaignType, err
}
