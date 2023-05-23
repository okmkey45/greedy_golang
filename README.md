# Greedy algorithms with Go

## Requirements

- Go 1.20

### Recommendations

If you have Docker installed and you're using VS Code, one good way is installing the [Dev Containers extension](https://code.visualstudio.com/docs/devcontainers/containers). By using this, you will avoid installations in your laptop and a container will be created for you with all necessary requirements

## Usage

This script computes the minimum number of coins needed to give back as change by using a greedy algorithm.

```sh
go run main.go
```

The above command will ask for two inputs: **Coin list** and **Change**

### Coin list

- It must be a list of integers separated by comma with no spaces. Example: 1,5,10
- Current implementation assumes that coin list is **sorted ascending**
- No checks have been implemented to verify that coin list is sorted ascending properly
- If you want to use the default coin list, just type `0` and hit enter. Default coin list: 1,5,10,25,50

### Change

- Only one integer value

Here's a full example of what you should see:

```txt
Enter coin list separated by comma
For default coin list ([1 5 10 25 50]), please enter '0'
Note: coin values must be sorted ascending. If not, unexpected behavior might occur
0

Enter change (only integer values)
63

Considering coin list: [1 5 10 25 50]
And change equals to: 63
The minimum number of coins for change are: 5
```

## Test

Just run:

```sh
go test ./...
```
