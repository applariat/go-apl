go-apl tests
============

The tests in this directory are only run manually.

integration
-----------

These tests will exercise the entire go-apl library. They will connect to the specified live API and make real changes. It is advisable to use a test account.

Run tests using:

    APL_API=XXX APL_SVC_USERNAME=XXX APL_SVC_PASSWORD=XXX go test -v -tags=integration ./integration

These variables can also be set via a $HOME/.apl/config.toml.
