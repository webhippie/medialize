# Changelog for 1.0.1

The following sections list the changes for 1.0.1.

## Summary

 * Fix #17: Enabled `rename` could result in error

## Details

 * Bugfix #17: Enabled `rename` could result in error

   When we enabled `rename` flag it could result in various errors related to cross-device link
   messages, we avoid that by falling back to a copy and delete function if this error appears.

   https://github.com/webhippie/medialize/issues/17


# Changelog for 1.0.0

The following sections list the changes for 1.0.0.

## Summary

 * Chg #7: Initial release of basic version

## Details

 * Change #7: Initial release of basic version

   Just prepared an initial basic version which could be released to the public.

   https://github.com/webhippie/medialize/issues/7


