package mongodb

import "go.temporal.io/server/common/metrics"

// Metrics of the MongoDB persistence plugin. They are defined here rather than
// in common/metrics so that the plugin does not modify upstream files; the
// metrics registry accepts definitions from any package.
var (
	persistenceMongoTransactionsStarted = metrics.NewCounterDef("persistence_mongo_transactions_started")
	persistenceMongoTransactionsFailed  = metrics.NewCounterDef("persistence_mongo_transactions_failed")
	persistenceMongoSessionsInProgress  = metrics.NewGaugeDef("persistence_mongo_sessions_in_progress")
	persistenceMongoMaxPoolSize         = metrics.NewGaugeDef("persistence_mongo_max_pool_size")
)
