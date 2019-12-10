package meilisearch

// ApiIndexes, index is an entity, like a table in SQL, with a specific schema definition. It gathers a collection of documents
// with the structure defined by the schema.
// An index is defined by an unique identifier uid that is generated by MeiliSearch (if none is given) on index
// creation. It also has a name to help you track your different indexes.
//
// Documentation: https://docs.meilisearch.com/references/indexes.html
type ApiIndexes interface {

	// Get the index relative information.
	Get(uid string) (*Index, error)

	// List all indexes.
	List() ([]Index, error)

	// Create an index.
	// The schema definition is optionally send through the body. If no schema has been defined when the first document
	// is sent it will be inferred based on that document.
	// If no UID is specified in the request a randomly generated UID will be returned.
	// It's associated to the new index. This UID will be essential to make all request over the created index.
	Create(request CreateIndexRequest) (*CreateIndexResponse, error)

	// Update an index name.
	Update(uid string, name string) (*Index, error)

	// Delete an index.
	Delete(uid string) (bool, error)

	// GetRawSchema get the schema of one index using the RawSchema type.
	GetRawSchema(uid string) (*RawSchema, error)

	// GetSchema get the schema of one index using the Schema type.
	GetSchema(uid string) (*Schema, error)

	// UpdateSchema an index schema by using the Schema type.
	UpdateSchema(uid string, schema Schema) (*UpdateIdResponse, error)

	// UpdateWithRawSchema an index schema by using the RawSchema type.
	UpdateWithRawSchema(uid string, schema RawSchema) (*UpdateIdResponse, error)
}

// ApiDocuments are objects composed of fields containing any data.
//
// Documentation: https://docs.meilisearch.com/references/documents.html
type ApiDocuments interface {

	// Get one document using its unique identifier.
	// documentPtr should be a pointer.
	Get(identifier string, documentPtr interface{}) error

	// Delete one document based on its unique identifier.
	Delete(identifier string) (*UpdateIdResponse, error)

	// Delete a selection of documents based on array of identifiers.
	Deletes(identifier []string) (*UpdateIdResponse, error)

	// List the documents in an unordered way.
	// WARNING: this route is a non-optimized route, it can be a little bit slow to answer.
	List(request ListDocumentsRequest, documentsPtr interface{}) error

	// AddOrUpdate a list of documents, update them if they already exist based on their unique identifiers.
	AddOrUpdate(documentsPtr interface{}) (*UpdateIdResponse, error)

	// Delete all documents in the specified index.
	ClearAllDocuments() (*UpdateIdResponse, error)
}

// Search through documents list in an index.
type ApiSearch interface {

	// Search for documents matching a specific query in the given index.
	Search(params SearchRequest) (*SearchResponse, error)
}

type ApiSynonyms interface {
	List(word string) ([]string, error)
	ListAll() ([]ListSynonymsResponse, error)
	Create(word string, synonyms []string) (*UpdateIdResponse, error)
	Update(word string, synonyms []string) (*UpdateIdResponse, error)
	Delete(word string) (*UpdateIdResponse, error)
	BatchCreate(request BatchCreateSynonymsRequest) (*UpdateIdResponse, error)
	DeleteAll() (*UpdateIdResponse, error)
}

type ApiStopWords interface {
	List() ([]string, error)
	Add(words []string) ([]UpdateIdResponse, error)
	Deletes(words []string) ([]UpdateIdResponse, error)
}

type ApiUpdates interface {
	Get(id int64) (*Unknown, error)
	List() ([]Unknown, error)
}

type ApiKey interface {
	Get(key string) (*APIKey, error)
	List() ([]APIKey, error)
	Create(request CreateApiKeyRequest) (*APIKey, error)
	Update(request UpdateApiKeyRequest) (*APIKey, error)
	Delete(key string) error
}

type ApiSettings interface {
	Get() (*Settings, error)
	AddOrUpdate(request Settings) (*UpdateIdResponse, error)
}

type ApiStats interface {
	Get() (*Stats, error)
	List() ([]Stats, error)
}

type ApiHealth interface {
	Get() error
	Set(health bool) error
}

type ApiVersion interface {
	Get() (*Version, error)
}

type ApiSystemInformation interface {
	Get() (*SystemInformation, error)
	GetPretty() (*SystemInformationPretty, error)
}