# FileStorage
This program receives exchange rates from the https://openexchangerates.org 
website and enters euro and dollar rates in rubles into the file currency.txt.

openexchangerates.org only allows you to send requests to its API 1000 times a month. 
Therefore, the program receives currency values every hour.

__By default,  file is stored in the data directory.__

__You may change the `volumes` section in the `docker-compose.yaml` file.__
## Getting started
1. Clone this repository
2. Add api key to `config.env` file
3. Run `docker-compose up -d --build` to start the project

## Environment variables
| Name     | Description               | Example                             |
|----------|---------------------------|-------------------------------------|
| `APP_ID` | openexchangerates api key | `N8DKXqKPP9t7ptmmFNxWireHpmtUxK8e`  |

### Example output:
```
€=102.77 $=91.74
```