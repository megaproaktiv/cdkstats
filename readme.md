# cdkstat

see Status of cdk stacks

1) Write CDK Stack names in csv

`cdk ls > stacks.csv`

2) `cdkstat`

You will see the Cloudformation State of *only* the Stacks managed by CDK

```txt
Name                             Status                           Description
----                             ------                           -----------
application                      CREATE_COMPLETE                  Application Server
securitygroups                   CREATE_COMPLETE                  Security Groups
lambda                           CREATE_COMPLETE                  Serverless Stuff
```

3) `ckdstat securitygroups`

```txt
Logical ID                       Pysical ID                       Type                             Status
----------                       ----------                       -----------                      -----------
SG1                              sg-00026dad358d5e2a3             AWS::EC2::SecurityGroup          CREATE_COMPLETE
CDKMetadata                      ae25d3b0-ce8c-11eb-a6c0-0ae8c75  AWS::CDK::Metadata               CREATE_COMPLETE
SG007                            sg-02393c66de6a16973             AWS::EC2::SecurityGroup          CREATE_COMPLETE
```


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
NoThereStack
```



`cdkstat` will give

```txt
Name                             Status
----                             ------
SecurityStack                    CREATE_COMPLETE
NoThereStack                     LOCAL_ONLY
```
