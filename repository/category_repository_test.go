package repository

import (
	"database/sql"
	"errors"
	"final-project-enigma-clean/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CategoryRepositoryTest struct {
	suite.Suite
	repo CategoryRepository
	mockDB *sql.DB
	mockSQL sqlmock.Sqlmock
}

func (suite *CategoryRepositoryTest) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSQL = mock
	suite.repo = NewCategoryRepository(suite.mockDB)
}

func TestCategoryRepositoryTest(t *testing.T)  {
	suite.Run(t, new(CategoryRepositoryTest))
}

func (suite *CategoryRepositoryTest) TestCreate_Success() {
	mockData := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.mockSQL.ExpectExec("INSERT INTO category").WithArgs(mockData.Id, mockData.Name).WillReturnResult(sqlmock.NewResult(1,1))
	err := suite.repo.Save(mockData)
	assert.NoError(suite.T(), err)
}

func (suite *CategoryRepositoryTest) TestCreate_Failed() {
	mockData := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.mockSQL.ExpectExec("INSERT INTO category").WithArgs(mockData.Id, mockData.Name).WillReturnError(errors.New("failed save category"))
	err := suite.repo.Save(mockData)
	assert.Error(suite.T(), err)
}

func (suite *CategoryRepositoryTest) TestFindAll_Success() {
	assets := []model.Category{
		{
			Id:   "1",
			Name: "Bergerak",
		},
	}
	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, v := range assets {
		rows.AddRow(v.Id, v.Name)
	}
	
	suite.mockSQL.ExpectQuery("SELECT id, name FROM category").WillReturnRows(rows)
	got, err := suite.repo.FindAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), assets, got)
}


func (suite *CategoryRepositoryTest) TestFindAll_FailedRows() {
	assets := []model.Category{
		{
			Id:   "1",
			Name: "Bergerak",
		},
	}
	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, v := range assets {
		rows.AddRow(v.Id, v.Name).RowError(0, errors.New("errors new scan"))
	}
	
	suite.mockSQL.ExpectQuery("SELECT id, name FROM category").WillReturnRows(rows)
	got, err := suite.repo.FindAll()
	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), assets, got)
}

func (suite *CategoryRepositoryTest) TestFindAll_Failed() {

	
	suite.mockSQL.ExpectQuery("SELECT id, name FROM category").WillReturnError(errors.New("failed get categories"))
	got, err := suite.repo.FindAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), got)
}

func (suite *CategoryRepositoryTest) TestFindById_Success() {
	assets := model.Category{
			Id:   "1",
			Name: "Bergerak",
		}
	row := sqlmock.NewRows([]string{"id", "name"}).AddRow(assets.Id, assets.Name)
	suite.mockSQL.ExpectQuery("SELECT id, name FROM category").WithArgs("1").WillReturnRows(row)
	got, err := suite.repo.FindById("1")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), assets, got)
}

func (suite *CategoryRepositoryTest) TestFindById_Failed() {
	suite.mockSQL.ExpectQuery("SELECT id, name FROM category").WithArgs("1").WillReturnError(errors.New("failed get category"))
	got, err := suite.repo.FindById("1")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Category{}, got)
}

func (suite *CategoryRepositoryTest) TestUpdate_Success() {
	mockData := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.mockSQL.ExpectExec("UPDATE category SET").WithArgs(mockData.Id, mockData.Name).WillReturnResult(sqlmock.NewResult(1,1))
	err := suite.repo.Update(mockData)
	assert.NoError(suite.T(), err)
}	

func (suite *CategoryRepositoryTest) TestUpdate_Failed() {
	mockData := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.mockSQL.ExpectExec("UPDATE category SET").WithArgs(mockData.Id, mockData.Name).WillReturnError(errors.New("failed update category"))
	err := suite.repo.Update(mockData)
	assert.Error(suite.T(), err)
}

func (suite *CategoryRepositoryTest) TestDelete_Success() {
	suite.mockSQL.ExpectExec("DELETE FROM category").WithArgs("1").WillReturnResult(sqlmock.NewResult(1, 1))
	gotErr := suite.repo.Delete("1")
	assert.NoError(suite.T(), gotErr)
}

func (suite *CategoryRepositoryTest) TestDelete_Failed() {
	suite.mockSQL.ExpectExec("DELETE FROM category").WithArgs("1").WillReturnError(errors.New("failed delete category"))
	gotErr := suite.repo.Delete("1")
	assert.Error(suite.T(), gotErr)
}