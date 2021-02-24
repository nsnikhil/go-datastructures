## ERX

#### Description
Erx provides a new way to create errors in go, embed kind, severity, operation and cause of the error in a single object.
You can stack erx.

---

#### Usage:

##### New vanilla erx
```
erx.WithArgs(errors.New("some error"))
```

##### Error with Kind
```
prebuilt kind
erx.WithArgs(errors.New("some error"), erx.ValidationError)

custom
erx.WithArgs(errors.New("some error"), erx.Kind("DependencyError"))
```

##### Error with Operation
```
erx.WithArgs(errors.New("some error"), erx.Operation("DB.FetchRecord"))
```

##### Error with Severity
```
prebuilt severity
erx.WithArgs(errors.New("some error"), erx.SeverityError)

custom
erx.WithArgs(errors.New("some error"), erx.Severity("Low"))
```

##### Combine All
```
erx.WithArgs(errors.New("some error"), erx.SeverityError, erx.ValidationError, erx.Operation("Client.Validate"))
```

##### Nested
```
erx.WithArgs(erx.Operation("Handler.UpdateUser"), erx.WithArgs(erx.Operation("Service.Update"), errors.New("some error"))
```

---

#### Api

##### Error()
Returns the cause of error of top most error in stack
```
err := erx.WithArgs(errors.New("some error"))
err.Error() // "some error"
```

##### Kind()
Returns the first non-empty kind in error stack
```
flat erx
erx.WithArgs(errors.New("some error"), erx.ValidationError)
err.Kind() // erx.ValidationError

nested erx
erx.WithArgs(erx.Operation("Handler.UpdateUser"), erx.WithArgs(erx.ValidationError, errors.New("some error"))
err.Kind() // erx.ValidationError
```

##### Operations()
Returns all the operations in stack
```
flat erx
erx.WithArgs(errors.New("some error"), erx.Operation("Handler.Update"))
err.Operations() // {"Handler.Update"}

nested erx
erx.WithArgs(erx.Operation("Handler.Update"), erx.WithArgs(erx.Operation("Service.Update"), errors.New("some error"))
err.Operations() // {"Handler.Update", "Service.Update"}
```

##### String()
Returns json representation of entire error stack
```
erx.WithArgs(erx.Operation("Handler.Update"), erx.Severity("Low"), erx.WithArgs(erx.Kind("validationError"), erx.Severity("Med"), errors.New("some error"))
err.String() // {"operation":"Handler.Update","severity":"Low","cause":{"kind":"validationError","severity":"Med","cause":"some error"}}
```

---

### Usage

1. `go get -u github.com/nsnikhil/erx`

---

### Contributing

1. Fork it (<https://github.com/nsnikhil/erx>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

---

### License

Copyright 2021 Nikhil Soni

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

