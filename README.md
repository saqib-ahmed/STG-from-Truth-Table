# State-Transition-Graph-from-CSV
This project is aimed to design a state transition graph of a moore finite state machine from a csv truth table containing next state values for each input.

## What does it do?
The go script in this project converts a csv formatted truth table in a dot script which is then read by the graphviz to generate the state transition graph of a moore finite state machine.

## Usage

### Pre-requisites:
Download and install [go-lang](https://golang.org/dl/) and [graphviz](http://www.graphviz.org/Download..php). That's all it requires.

### CSV format:
The csv should be formatted with "empty cell or number","Present state", "Input bits", "empty value", "next state", "output". Here, present state, input bits, next state and output should have values and should have same offset from start. Other two fields may or may not have any sort of values. 
> Currently this script takes out csv values from these offsets. In subsequent releases, I'll try to read values by name of heading. That makes more sense.

For example, [truth_fsm.csv](https://github.com/saqibahmed515/State-Transition-Graph-from-CSV/blob/master/truth_fsm.csv) is a valid csv file.  The heading names are irrelevant and ignored by default. Just give the values in this order.

### Generating STG:
General command format to execute the script is:
```bash
go run script.go </path/to/csvFile> <dpi resolution>
```
simply execute this command in the project directory with a valid csv file.
#### Example
In this project, you can execute:
```bash
go run script.go truth_fsm.csv 300
```
#### Sample Output
Sample outputs with various dpi resolutions are provided in [Sample Pictures](https://github.com/saqibahmed515/State-Transition-Graph-from-CSV/tree/master/sample%20pictures) folder.
![Sample image](https://github.com/saqibahmed515/State-Transition-Graph-from-CSV/blob/master/sample%20pictures/out_150dpi.png)


## Sample project
A sample project is also included which has complete flow of a solution.

### Problem statement:
> 1) A vending machine can accept a nickle (5 cents), dime (10 cents) or a quarter (25 cents).
> 2) A vending machine Will dispense two types of candy. One type Will cost 35 cents, the other will only cost 30 cents. There will be an input signal linked to a button that indicates which type of candy is desired.
> 3) When enough money has been entered to buy the chosen item, the vending machine Will dispense the appropriate piece of candy, but only as much as one of each. This means that overshooting by 5, 10 or 15 cents will give exact amount of change but overshooting 20 cents will give 15 cents back only.

### Solution flow:
* The truth table for the next state and output has been made in excel file given at [truth_fsm.xlsx](https://github.com/saqibahmed515/State-Transition-Graph-from-CSV/blob/master/Sample%20Project/truth_fsm.xlsx) which is then exported as a csv [truth_fsm.csv](https://github.com/saqibahmed515/State-Transition-Graph-from-CSV/blob/master/truth_fsm.csv). This is created with the problem requirements in mind.

* State transition graph was then created with the script provided in the project.
* [Logic Friday](http://sontrak.com/downloads.html) is then used to create truth tables for combinational circuits. All the logic friday files are also given in the [sample project](https://github.com/saqibahmed515/State-Transition-Graph-from-CSV/tree/master/Sample%20Project).
* To reduce the input bits, shannon's expansion has been used with 4 `8x1` multiplexers for four output bits.
* Each `*.lfcn` file contains truth tables for each value corresponding to the input of `8x1` multiplexers. Logic friday automatically creates the minimized equations and combinational circuits from the truth tables.
* See the report [final.pdf](https://github.com/saqibahmed515/State-Transition-Graph-from-CSV/blob/master/Sample%20Project/Final%20Project.pdf) for more details on solving the project and to see the actual outputs.