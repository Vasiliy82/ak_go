package exchange

import (
	"database/sql"

	"gopkg.in/guregu/null.v3"
)

type Method string
type Action string
type Property string

type Filter struct {
	Name   string   `json:"Name"`
	Null   *bool    `json:"Null"`
	Value  *string  `json:"Value"`
	Values []string `json:"Values"`
}

type Result interface{}

type Results struct {
	Error   *string  `json:"Error"`
	Results []Result `json:"Results"`
}

type Params struct {
	LifeTime  null.String `json:"LifeTime"`
	Session   *int32      `json:"Session"`
	ReNew     *bool       `json:"ReNew"`
	FixErrors *bool       `json:"FixErrors"`
	End       *bool       `json:"End"`
	Target    *string     `json:"Target"`
	CopyDir   *string     `json:"CopyDir"`
	Limit     *int64      `json:"Limit"`
	Offset    *int64      `json:"Offset"`
	Fields    *string     `json:"Fields"`
	Filters   []Filter    `json:"Filters"`
	Results   *Results    `json:"Results"`
}

type Session struct {
	Session int    `json:"Session"`
	Actions string `json:"Actions"`
	Count   int    `json:"Count"`
	Primary string `json:"Primary"`
}

const (
	MethodOnline  Method = "ONLINE"
	MethodOffline Method = "OFFLINE"
)

const (
	actionAccess           Action = "ACCESS"
	actionObjects          Action = "OBJECTS"
	actionTargets          Action = "TARGETS"
	actionMetadata         Action = "METADATA"
	actionColumns          Action = "COLUMNS"
	actionJSONColumns      Action = "JSON:COLUMNS"
	actionXMLColumns       Action = "XML:COLUMNS"
	actionChecksum         Action = "CHECKSUM"
	actionJSONChecksum     Action = "JSON:CHECKSUM"
	actionXMLChecksum      Action = "XML:CHECKSUM"
	actionBegin            Action = "BEGIN"
	actionRows             Action = "ROWS"
	actionTableRows        Action = "TABLE:ROWS"
	actionBulkRows         Action = "BULK:ROWS"
	actionRowsDeleted      Action = "ROWS:DELETED"
	actionTableRowsDeleted Action = "TABLE:ROWS:DELETED"
	actionBulkRowsDeleted  Action = "BULK:ROWS:DELETED"
	actionRefill           Action = "REFILL"
	actionDataset          Action = "DATASET"
	actionTableDataset     Action = "TABLE:DATASET"
	actionBulk             Action = "BULK"
	actionXMLSelect        Action = "XML:SELECT"
	actionXMLSelectText    Action = "XML:SELECT:TEXT"
	actionXMLOutput        Action = "XML:OUTPUT"
	actionJsonSelect       Action = "JSON:SELECT"
	actionJsonOutput       Action = "JSON:OUTPUT"
	actionResults          Action = "RESULTS"
	actionTableResults     Action = "TABLE:RESULTS"
	actionEnd              Action = "END"
)
const (
	PropertyTarget    Property = "Target"
	PropertyView      Property = "View"
	PropertyProcedure Property = "Procedure"
	PropertyFunction  Property = "Function"
	PropertyTempDir   Property = "TempDir"
	PropertyPrimary   Property = "Primary"
	PropertyParent    Property = "Parent"
	PropertyMapping   Property = "Mapping"
	PropertyFields    Property = "Fields"
	PropertyFilters   Property = "Filters"
	PropertyOptions   Property = "Options"
)

type Object struct {
	Name   string
	Method Method
}

type Client struct {
	db         *sql.DB
	conn       *sql.Conn
	subscriber string
	object     Object
}

type Rows interface{}

func NewClient(db *sql.DB, subscriber string, object Object) *Client {
	return &Client{
		db:         db,
		conn:       nil,
		subscriber: subscriber,
		object:     object,
	}
}

func (c *Client) Objects() ([]Object, error) {
	_ = actionObjects
	return nil, nil
}

func (c *Client) WaitForChanges() (bool, error) {
	return false, nil
}

func (c *Client) Begin(params *Params) (*Session, error) {
	_ = actionBegin
	return &Session{}, nil
}

func (c *Client) Refill(params *Params) error {
	_ = actionRefill
	return nil
}

func (c *Client) Rows(params *Params, target Rows) error {
	_ = actionRows
	return nil
}

func (c *Client) RowsDeleted(params *Params, target Rows) error {
	_ = actionRows
	return nil
}

func (c *Client) Dataset(params *Params, target Rows) error {
	_ = actionDataset
	return nil
}

func (c *Client) Results(Params *Params) error {
	return nil
}

func (c *Client) End() error {
	return nil
}

func Test() {
	db, _ := sql.Open("", "")
	c := NewClient(db, "", Object{Name: "", Method: MethodOnline})
	c.Begin(nil)
	c.Begin(&Params{LifeTime: null.NewString("00:10:00", true)})
	_ = c

}
