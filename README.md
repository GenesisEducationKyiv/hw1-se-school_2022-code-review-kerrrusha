# BTC-API Description

This is an API service written in Golang. The service provides the ability to receive the current bitcoin rate in hryvnia equivalent, sign new mails to receive up-to-date information about the bitcoin rate, as well as the ability to send each signed mail a letter with the current bitcoin rate.

To implement the function of obtaining the current bitcoin rate, the service "coinapi.io" was chosen. This service provides free access to cryptocurrency rates, which includes a limit of 100 daily requests. The main packages in developing this feature are "net/http" and "encoding/json".

Signed mail data is stored locally in a ".json" file. A SMTP Google account was created to send letters.
The built-in package "net/smtp" was used.

The whole solution was implemented using only the packages built into the Golang language, so that at the stage of preparation for dockerizing the solution there was no need to create a go.sum file, which is designed to verify the hash sums of downloaded third-party packages.

All code has been split into separate sub-packages as part of a modular approach to writing programs. Although at this stage the service may look too small for this architectural approach, future expansion of the project may come in handy. Also, frequently performed functions were placed in general packages (for example, error handling or sending a response from the server), and all structures were placed in the "models" folder. This will make it easier for other developers to read the solution.

Working constants, such as the URL of the service that provides the bitcoin rate, as well as the name of the file that stores signed email addresses, are stored in the "config" package.

The server is running at address 0.0.0.0 (this address worked great when testing the server in Docker), the expected port is 8000, although you can define it yourself by creating an environment variable.

# Logic

1. GET /rate
The server will try to get the json response from the aforementioned third-party API that provides the bitcoin rate. Since we can expect an error from the server (for example, the connection key is incorrect/outdated), we process two options:
1) received an error message - return 400 code and error text
2) received a response containing information about the bitcoin rate - turn the actual value of the bitcoin rate into an integer, and return the response 200 with the received value.

2. POST/subscribe
Here the server will have to work with the .json file that stores the mail. If it is missing, a new file with the "list of mail" structure will be created. The main method responsible for processing this request, using the indexOfEmail(filename, email) method, checks if the specified mail already exists. By the way, this function checks emails regardless of case, because as a rule, the spelling of an email address itself does not depend on it. So, we have two options for the development of events:
1) indexOfEmail(filename, e-mail) returned a value other than -1 (the index of the first occurrence of the mail in the list of mails in the file) - this means that such mail already exists. return 409 code indicating an error.
2) indexOfEmail(filename, email) returned -1 - using the writeNewEmailToFile(filename, email) method, save the sent mail to a .json file, return 200 success code.

3. POST /send email
The processing of this request begins with getting the current bitcoin rate using the GetBitcoinPriceUAH() method used in the /rate stage. As we know, when executing a request to get a bitcoin rate, a connection error may occur, so we expect two responses from this method:
1) error - immediately return 400 error code
2) the bitcoin rate - then we read all signed mails from the file using the readEmails(filename) method, and each of the mails is separately sent using the Gmail SMTP server with the received bitcoin rate. As a result of the execution, we send the success code 200.
