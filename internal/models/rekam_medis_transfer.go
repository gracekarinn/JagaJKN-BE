package models

import (
	"time"
)

type RekamMedisTransfer struct {
    ID                string           `gorm:"type:varchar(50);primaryKey" json:"id"`
    NoSEP             string           `gorm:"type:varchar(20)" json:"noSEP"`
    RecordKesehatan   *RecordKesehatan `gorm:"foreignKey:NoSEP;references:NoSEP" json:"recordKesehatan"`
    SourceFaskes      string           `gorm:"type:varchar(20)" json:"sourceFaskes"`
    DestinationFaskes string           `gorm:"type:varchar(20)" json:"destinationFaskes"`
    TransferReason    string           `gorm:"type:text" json:"transferReason"`
    TransferTime      time.Time        `json:"transferTime"`
    Status           string           `gorm:"type:status_transfer" json:"status"`
    CreatedAt        time.Time        `json:"createdAt"`
    UpdatedAt        time.Time        `json:"updatedAt"`
}
func (t *RekamMedisTransfer) ToJSON() map[string]interface{} {
    return map[string]interface{}{
        "id":                t.ID,
        "noSEP":            t.NoSEP,
        "recordKesehatan":  t.RecordKesehatan.ToBlockchainRecord(),
        "sourceFaskes":     t.SourceFaskes,
        "destinationFaskes": t.DestinationFaskes,
        "transferReason":    t.TransferReason,
        "transferTime":      t.TransferTime,
        "status":           t.Status,
        "createdAt":        t.CreatedAt,
        "updatedAt":        t.UpdatedAt,
        "user":             t.RecordKesehatan.User.ToJSON(), 
    }
}