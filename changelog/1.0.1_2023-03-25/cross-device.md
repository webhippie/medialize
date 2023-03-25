Bugfix: Enabled `rename` could result in error

When we enabled `rename` flag it could result in various errors related to
cross-device link messages, we avoid that by falling back to a copy and delete
function if this error appears.

https://github.com/webhippie/medialize/issues/17
