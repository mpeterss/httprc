# httprc

`httprc` is a HTTP "Refresh" Cache. Its aim is to cache a remote resource that
can be fetched via HTTP, but keep the cached content up-to-date based on periodic
refreshing.

# SYNOPSIS

<!-- INCLUDE(httprc_example_test.go) -->
<!-- END INCLUDE -->

# Sequence Diagram

```mermaid
sequenceDiagram
  autonumber
  actor User
  participant httprc.Cache
  participant httprc.Storage
  User->>httprc.Cache: Fetch URL `u`
  activate httprc.Storage
  httprc.Cache->>httprc.Storage: Fetch local cache for `u`
  alt Cache exists
    httprc.Storage-->httprc.Cache: Return local cache
    httprc.Cache-->>User: Return data
    Note over httprc.Storage: If the cache exists, there's nothing more to do.<br />The cached content will be updated periodically in httprc.Refresher
    deactivate httprc.Storage
  else Cache does not exist
    activate httprc.Fetcher
    httprc.Cache->>httprc.Fetcher: Fetch remote resource `u`
    httprc.Fetcher-->>httprc.Cache: Return fetched data
    deactivate httprc.Fetcher
    httprc.Cache-->>User: Return data
    httprc.Cache-)httprc.Refresher: Enqueue into auto-refresh queue
    activate httprc.Refresher
    loop Refresh Loop
      Note over httprc.Storage,httprc.Fetcher: Cached contents are updated synchronously
      httprc.Refresher->>httprc.Refresher: Wait until next refresh
      httprc.Refresher-->>httprc.Fetcher: Request fetch
      httprc.Fetcher->>httprc.Refresher: Return fetched data
      httprc.Refresher-->>httprc.Storage: Store new version in cache
      httprc.Refresher->>httprc.Refresher: Enqueue into auto-refresh queue (again)
    end
    deactivate httprc.Refresher
  end
```
