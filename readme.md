# cdkstat

see Status of cdk stacks

1) Write CDK Stack names in csv

`cdk ls > stacks.csv`

2) `cdkstat`

You will see the Cloudformation State of *only* the Stacks managed by CDK

## Installation

1. Create binary

`go build -o cdkstat main.go`

2. Move binary into PATH

e.g.

`mv cdkstat /usr/local/bin/`

## Example output

If you have those stacks locally:

```bash
cat stacks.csv
SecurityStack
PortalServerStack
NoThereStack
```



`cdkstat` will give

```txt
Name                             Status
----                             ------
PortalServerStack                CREATE_COMPLETE
SecurityStack                    CREATE_COMPLETE
NoThereStack                     LOCAL_ONLY
```
