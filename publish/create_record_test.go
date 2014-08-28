package publish_test

import (
	_ "github.com/mattn/go-sqlite3"

	"testing"
)

func TestCreateStructFromDraft(t *testing.T) {
	name := "create_product_from_draft"
	pbdraft.Create(&Product{Name: name, Color: Color{Name: name}})

	if !pbprod.First(&Product{}, "name = ?", name).RecordNotFound() {
		t.Errorf("record should not be found in production db")
	}

	if pbdraft.First(&Product{}, "name = ?", name).RecordNotFound() {
		t.Errorf("record should be found in draft db")
	}

	if pb.Table("colors").First(&Color{}, "name = ?", name).Error != nil {
		t.Errorf("color should be saved")
	}

	if pb.Table("colors_draft").First(&Color{}, "name = ?", name).Error == nil {
		t.Errorf("no colors_draft table")
	}

	var product Product
	pbdraft.First(&product, "name = ?", name)
	if pbdraft.Model(&product).Related(&product.Color); product.Color.Name != name {
		t.Errorf("should be able to find related struct")
	}
}

func TestCreateStructFromProduction(t *testing.T) {
	name := "create_product_from_production"
	pbprod.Create(&Product{Name: name, Color: Color{Name: name}})

	if pbprod.First(&Product{}, "name = ?", name).RecordNotFound() {
		t.Errorf("record should not be found in production db")
	}

	if pbdraft.First(&Product{}, "name = ?", name).RecordNotFound() {
		t.Errorf("record should be found in draft db")
	}

	if pb.Table("colors").First(&Color{}, "name = ?", name).Error != nil {
		t.Errorf("color should be saved")
	}

	var product Product
	pbprod.First(&product, "name = ?", name)
	if pbprod.Model(&product).Related(&product.Color); product.Color.Name != name {
		t.Errorf("should be able to find related struct")
	}
}
