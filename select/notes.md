testing HTTP is so common Go has tools in the standard lib to help us
- net.http/httptest
- we don't want to rely on external service to test our code

select
- construct that lets us synchronise processes
- myVar := <-ch lets you wait for values to be sent to channels - which is blocking
- select allows you to wait on multiple channels.  The first one to send a value 'wins' and the code underneath is executed
- Ofter you'll want to include time.After on one of the cases, to prevent infinite blocking
