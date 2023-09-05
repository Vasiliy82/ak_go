package exchange

import (
	"database/sql"
)

type AnyString string

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

type ParamsBegin struct {
	LifeTime  *string `json:"LifeTime"`
	ReNew     *bool   `json:"ReNew"`
	FixErrors *bool   `json:"FixErrors"`
	End       *bool   `json:"End"`
	// Target    *string  `json:"Target"`
}

type ParamsRefill struct {
	Session *int32 `json:"Session"`
	ReNew   *bool  `json:"ReNew"`
	//	Target    *string  `json:"Target"`
	Limit  *int32 `json:"Limit"`
	Offset *int64 `json:"Offset"`
}

type ParamsRows struct {
	Session *int32 `json:"Session"`
	//	Target    *string  `json:"Target"`
	Limit   *int32   `json:"Limit"`
	Offset  *int64   `json:"Offset"`
	Filters []Filter `json:"Filters,omitempty"`
}

type ParamsDataset struct {
	Session *int32   `json:"Session"`
	Target  *string  `json:"Target"`
	Limit   *int32   `json:"Limit"`
	Offset  *int64   `json:"Offset"`
	Fields  *string  `json:"Fields"`
	Filters []Filter `json:"Filters,omitempty"`
}

type ParamsResult struct {
	Session *int32 `json:"Session"`
	//	Target    *string  `json:"Target"`
	Limit   *int32   `json:"Limit"`
	Offset  *int64   `json:"Offset"`
	Results *Results `json:"Results"`
}

type ParamsEnd struct {
	Session *int32 `json:"Session"`
}

type Session struct {
	Session int32  `json:"Session"`
	Actions string `json:"Actions"`
	Count   int32  `json:"Count"`
	Primary string `json:"Primary"`
}

const (
	MethodOnline  Method = "ONLINE"
	MethodOffline Method = "OFFLINE"
)

const (
	actionObjects     Action = "OBJECTS"
	actionBegin       Action = "BEGIN"
	actionRows        Action = "ROWS"
	actionRowsDeleted Action = "ROWS:DELETED"
	actionRefill      Action = "REFILL"
	actionDataset     Action = "DATASET"
	actionResults     Action = "RESULTS"
	actionEnd         Action = "END"
	// пока не удалять, могут пригодиться в будущем
	// actionAccess           Action = "ACCESS"
	// actionTargets          Action = "TARGETS"
	// actionMetadata         Action = "METADATA"
	// actionColumns          Action = "COLUMNS"
	// actionJSONColumns      Action = "JSON:COLUMNS"
	// actionXMLColumns       Action = "XML:COLUMNS"
	// actionChecksum         Action = "CHECKSUM"
	// actionJSONChecksum     Action = "JSON:CHECKSUM"
	// actionXMLChecksum      Action = "XML:CHECKSUM"
	// actionTableRows        Action = "TABLE:ROWS"
	// actionBulkRows         Action = "BULK:ROWS"
	// actionTableRowsDeleted Action = "TABLE:ROWS:DELETED"
	// actionBulkRowsDeleted  Action = "BULK:ROWS:DELETED"
	// actionTableDataset     Action = "TABLE:DATASET"
	// actionBulk             Action = "BULK"
	// actionXMLSelect        Action = "XML:SELECT"
	// actionXMLSelectText    Action = "XML:SELECT:TEXT"
	// actionXMLOutput        Action = "XML:OUTPUT"
	// actionJsonSelect       Action = "JSON:SELECT"
	// actionJsonOutput       Action = "JSON:OUTPUT"
	// actionTableResults     Action = "TABLE:RESULTS"
)

type Object struct {
	Name   string
	Method Method
}

type Client struct {
	db         *sql.DB
	conn       *sql.Conn
	subscriber string
	object     *Object
}

type Rows interface{}

// Конструктор
// Аргументы:
// db *sql.DB - подключение к БД (будет использовано для порождения connect-а)
// subscriber string - подписчик
// object exchange.Object - объект экспорта
// Возвращаемый результат:
// ссылка на новый экземпляр объекта
func NewClient(db *sql.DB, subscriber string, object *Object) *Client {
	return &Client{
		db:         db,
		conn:       nil,
		subscriber: subscriber,
		object:     object,
	}
}

// Возвращает список объектов, доступных подписчику, указанному при создании объекта
// Возвращаемый результат:
// Массив объектов типа Exchange.Object
// Алгоритм работы:
// Создает новое подключение и выполняет:
// EXEC [Export].[Execute] @SubScriber=c.subscriber, @Action=actionOjects
// анализирует полученный DATASET, заполняет массив и возвращает его в качестве результата.
func (c *Client) Objects() ([]Object, error) {
	_ = actionObjects
	return nil, nil
}

// Блокирующий вызов, который прервется в случае, если в объекте экспорта будут обнаружены изменения.
// Аргументы:
// timeoutMS - суммарное время ожидания, рекомендуется не более 2 мин
// delayMS - периодичность проверки, рекомендуется не менее 0.5 сек
// Алгоритм работы:
// 1. проверить входные аргументы: timeoutMS > 0, delayMS > 0;
// 2. проверить допустимость операции:
//   - у объекта экспорта Method должен быть MethodOnline;
//   - отсутствует активная сессия
//
// 3. преобразовать входные параметры timeoutMS и delayMS в нужный формат и выполнить:
// EXEC [Export].[Wait For Changes] @Object=c.object.Name, @SubScriber=c.subscriber, @Timeout=timeoutMS, @Delay=delayMS
// переменные timeoutMS и delayMS передавать в MS SQL в виде строк в формате: чч:мм:сс.мс
func (c *Client) WaitForChanges(timeoutMS, delayMS int32) (bool, error) {
	return false, nil
}

// Начинает сессию обмена с Exchange
// Аргументы:
// params *ParamsBegin - параметры, будут переданы в Exchange
// Возвращаемый результат:
// ссылка на объект Session
// Алгоритм работы:
// 1. проверить допустимость операции:
//   - у объекта экспорта Method должен быть MethodOnline;
//   - отсутствует активная сессия;
//
// 2. выполнить
// EXEC [Export].[Execute] @Object=c.object.Name, @SubScriber=c.subscriber, @Action=actionBegin, @Params=params, @Output = @o OUT
// (params передается как строка в формате JSON)
// 3. в структуре c сохранить признак того, что сессия открыта
// 4. результат, возвращенный из параметра @Output, привести к типу Session, и вернуть в качестве результата
func (c *Client) Begin(params *ParamsBegin) (*Session, error) {
	_ = actionBegin
	return &Session{}, nil
}

// Выполняет операцию REFILL с объектом
// Аргументы:
// params *ParamsRefill - параметры, будут переданы в Exchange
// Алгоритм работы:
// 1. проверить допустимость операции: есть активная сессия;
// 2. выполнить
// EXEC [Export].[Execute] @Object=c.object.Name, @SubScriber=c.subscriber, @Action=actionRefill, @Params=params
// (params передается как строка в формате JSON)
func (c *Client) Refill(params *ParamsRefill) error {
	_ = actionRefill
	return nil
}

// Возвращает ключи изменившихся строк и характер изменений (I/U/R/D)
// Аргументы:
// params *ParamsRows - параметры, будут переданы в Exchange
// target Rows - указатель на слайс структур, куда будет записан результат
// Возвращаемый результат:
// Ключи измененных записей и характер изменений: (PK, @Action), где @Action может принимать значения: I/U/R/D.
// Результат записывается в Rows
// Алгоритм работы:
// 1. проверить допустимость операции: есть активная сессия;
// 2. выполнить
// EXEC [Export].[Execute] @Object=c.object.Name, @SubScriber=c.subscriber, @Action=actionRows, @Params=params
// (params передается как строка в формате JSON)
// 3. Возвращенный DATASET сохранить в Rows
// Примечание: т.к. методы Rows, RowsDeleted очень схожи, следует сделать приватный метод rows с доп.агрументом action.
func (c *Client) Rows(params *ParamsRows, target Rows) error {
	_ = actionRows
	return nil
}

// Возвращает ключи удаленных строк
// Аргументы:
// params *ParamsRows - параметры, будут переданы в Exchange
// target Rows - указатель на слайс структур, куда будет записан результат
// Возвращаемый результат:
// Ключи удаленных записей (PK), результат записывается в Rows
// Алгоритм работы:
// 1. проверить допустимость операции: есть активная сессия;
// 2. выполнить
// EXEC [Export].[Execute] @Object=c.object.Name, @SubScriber=c.subscriber, @Action=actionRows, @Params=params
// (params передается как строка в формате JSON)
// 3. Возвращенный DATASET сохранить в Rows
// Примечание: т.к. методы Rows, RowsDeleted очень схожи, следует сделать приватный метод rows с доп.агрументом action.
func (c *Client) RowsDeleted(params *ParamsRows, target Rows) error {
	_ = actionRowsDeleted
	return nil
}

// Возвращает изменившиеся строки
// Аргументы:
// params *ParamsDataset - параметры, будут переданы в Exchange
// Возвращаемый результат:
// изменившиеся строки, результат записывается в Rows
// Алгоритм работы:
// 1. проверить допустимость операции:
//   - если у объекта импорта Method = MethodOnline:
//   - должна быть активная сессия;
//   - если у объекта импорта Method = MethodOffline:
//   - в params должен быть указан параметр Target;
//
// 2. выполнить
// EXEC [Export].[Execute] @Object=c.object.Name, @SubScriber=c.subscriber, @Action=actionDataset, @Params=params
// (params передается как строка в формате JSON)
// 3. Возвращенный DATASET сохранить в Rows
func (c *Client) Dataset(params *ParamsDataset, target Rows) error {
	_ = actionDataset
	return nil
}

// Уведомляет Exchange об успешной обработке строк
// Аргументы:
// params *ParamsResults - параметры, будут переданы в Exchange
// Алгоритм работы:
// 1. проверить допустимость операции: есть активная сессия;
// 2. выполнить
// EXEC [Export].[Execute] @Object=c.object.Name, @SubScriber=c.subscriber, @Action=actionResults, @Params=params
// (params передается как строка в формате JSON)
func (c *Client) Results(params *ParamsResult) error {
	_ = actionResults
	return nil
}

// Завершает сессию
// Аргументы:
// params *ParamsEnd - параметры, будут переданы в Exchange
// Алгоритм работы:
// 1. проверить допустимость операции: есть активная сессия;
// 2. выполнить
// EXEC [Export].[Execute] @Object=c.object.Name, @SubScriber=c.subscriber, @Action=actionEnd, @Params=params
// (params передается как строка в формате JSON)
// 3. в структуре c сохранить признак того, что сессия закрыта
func (c *Client) End(params ParamsEnd) error {
	_ = actionEnd
	return nil
}

// для удобства заполнения "NULLABLE"-параметров в params
func Str2ref(s string) *string {
	result := s
	return &result
}
