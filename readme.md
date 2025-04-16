Train ticket price calculator

## Problem statement

We want to develop a module to help IRCTC to calculate the ticket price for a given train. Idea is to give this module as an extension to different apps in future. Below are some examples

## Example 1

A train number 12345 has only sleeper and general coaches. The train runs between Mumbai to Pune and it has stops as shown below

| Start - Mumbai | Stop 1 - Karjat | Stop 2 - Lonavala | Stop 3 - Chinchwad | Stop 4 - Pune |
|----------------|-----------------|-------------------|--------------------|---------------|

Passengers can board at any stop and leave at _any stop ahead._

**This train follows fix-per-station pricing strategy as below**

For each stop passenger has to pay

Rs. 20/- for general coach

Rs. 40/- for sleeper coach

### Cases

| \# of passengers | Train number | Coach   | Start station | End station | Total Price              |
|------------------|--------------|---------|---------------|-------------|--------------------------|
| 1                | 12345        | general | Karjat        | Chinchwad   | **40**/-                 |
| 3                | 12345        | sleeper | Mumbai        | Pune        | 3 \* 40 \* 4 = **480**/- |

## Example 2

A train number 12346 has only sleeper and general coaches. The train runs between Mumbai to Pune and it has stops as shown below

| Start - Mumbai | Stop 1 - Karjat | Stop 2 - Lonavala | Stop 3 - Chinchwad | Stop 4 - Pune |
|----------------|-----------------|-------------------|--------------------|---------------|

Passengers can board at any stop and leave at _any stop ahead._

**This train follows variable-between-station pricing strategy as below**

| Coach   | Start station | End station | Cost per seat per station |
|---------|---------------|-------------|---------------------------|
| general | Mumbai        | Lonavala    | 30/-                      |
| general | Lonavala      | Pune        | 20/-                      |
| sleeper | Mumbai        | Lonavala    | 50/-                      |
| sleeper | Lonavala      | Pune        | 40/-                      |

Rs. 20/- for general coach

Rs. 40/- for sleeper coach

### Cases

| \# of passengers | Train number | Coach   | Start station | End station | Total Price                                                                                                                                                  |
|------------------|--------------|---------|---------------|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1                | 12346        | general | Karjat        | Chinchwad   | Karjat to Lonavala 30/-<br><br>Lonavala to Chinchwad 20/-<br><br>Total = **50**/-                                                                            |
| 3                | 12346        | sleeper | Mumbai        | Pune        | Mumbai to Lonavala 30/- per passenger per station<br><br>Lonavala to Pune 20/- per passenger per station<br><br>Total = 3 \* (50 \* 2 + 40 \* 2) = **540**/- |

We need your help in creating this module.

**Note:**

1. REST apis are not expected.
2. Don’t use any frameworks e.g. spring, micronaut etc
3. Databases are not expected to be integrated

### Expectations

1. Follow clean code practices
2. Follow coding practices
3. Write tests
