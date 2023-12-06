# Error handling

## About it

- Error handling in GraphQL is very different from RESTful APIs, where developers can rely on HTTP status codes to handle errors. There are several types of errors. Like network errors, request errors, and field errors. Here I mainly discuss these two types of errors: request errors and field errors. 

## Request Errors

- It happens before execution starts. This is usually the requesting client's fault. We can divide request errors into two main types of errors: syntax errors and validation errors.
  - Syntax errors are errors like incorrectly formatted query.
  - Validation errors are like query contains non-existent schema field.

- If a request error is raised, execution does not begin and the data entry in the response must not be present. The errors entry must include the error.

## Field Errors

- It happens during execution from a particular field. This errors are typically the fault of GraphQL service. For example, if the request is valid but the server cannot find the data in the data source or the data source is not available, a field error is returned.
- If a field error is raised, execution attempts to continue and a partial result is produced. The data entry in the response must be present. The errors entry should include all raised field errors.

## Use Extensions to Make Error Readable

- Some GraphQL engines have the extensions field to store error messages and make these messages easier to read. Developers can also create their own error type for this purpose.


