# Metrics

| Name                                                | Help                                                                                                                                     | Type             |
| :-------------------------------------------------- | :--------------------------------------------------------------------------------------------------------------------------------------- | :--------------- |
| gremcos_cosmos_statuscode_total                     | Counts the number of responses from cosmos separated by status code.                                                                     | Labelled Counter |
| gremcos_cosmos_retry_after_ms                       | The time in milliseconds suggested by cosmos to wait before issuing the next query.                                                      | Histogram        |
| gremcos_cosmos_request_charge_per_query             | Cosmos DB reports a request charge accumulated for all responses of one query. This metric represents that value.                        | Gauge            |
| gremcos_cosmos_request_charge_per_queryresponse_avg | Cosmos DB reports a request charge each of the responses of one query. This metric represents the average of these values for one query. | Gauge            |
| gremcos_cosmos_request_charge_total                 | The accumulated request charge over all queries issued so far.                                                                           | Counter          |
| gremcos_cosmos_server_time_per_query_ms             | The time spent in ms for one query.                                                                                                      | Gauge            |
| gremcos_cosmos_server_time_per_queryresponse_avg_ms | The average time spent in ms for one query per response.                                                                                 | Gauge            |
| gremcos_cosmos_connectivity_errors_total            | The amount of errors happened when creating a new connection.                                                                            | Counter          |
| gremcos_cosmos_connection_usage_total               | The amount of reads, writes and pings that where made (the label is called kind). Errors that happened are labelled as error=true.       | Labelled Counter |
| gremcos_cosmos_request_errors_total                 | The accumulated number of request errors.                                                                                                | Counter          |
| gremcos_cosmos_request_retries_total                | The accumulated number of retried requests.                                                                                              | Counter          |
| gremcos_cosmos_request_retry_timeouts_total         | The accumulated number of timeouts that happened for request retries.                                                                    | Counter          |
