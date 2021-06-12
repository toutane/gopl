package cli

var HelpMessages = map[string]string{"help": help, "auth": auth, "read": read, "create": create, "close": closed}

const help = `
 Create, read, update and close GitHub issues from the command line.

 USAGE
   
   ./gitool <command> <subcommand> [flags]

 CORE COMMANDS

   read:       Read issue(s).
   create:     Create a new issue.
   update:     Update an issue.
   close:      Close an issue.

 ADDITIONAL COMMANDS

   auth:       Login, logout and view status of your authentication.
   `
const read = `
 USAGE

   ./gitool read [flags]

 FLAGS
   
   --repo       Select issue's repository (required).
   --username   Select issue's owner (by default logged user).
   --state      Select issue's state (open, closed, all).
   --number     Select issue's number (by default all issues of repo are listed).
   `

const create = `
 USAGE

   ./gitool create [flags]

 FLAGS
   
   --repo       Select new issue's repository (required).
   --title      New issue's title (required, by default: New issue by 'your GitHub username').
   `

const update = `
 USAGE

   ./gitool update [flags]

 FLAGS
   
   --repo       Select issue's repository (required).
   --number     Number of the issue you want to update (required).
   `

const closed = `
 USAGE

   ./gitool close [flags]

 FLAGS
   
   --repo       Select issue's repository (required).
   --number     Number of the issue you want to close (required)
   `
const auth = `
 USAGE
  
   ./gitool auth <command>

 CORE COMMANDS
                
   login:      Authenticate to a GitHub account.
   logout:     Log out of a GitHub account.
   status:     View authenticate status.
	`
