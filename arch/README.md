# arch

Generate HCL version of terraform configuration

```sh
$ cd $GOPATH/src/arch
$ make linux | darwin | windows | all
```

Generation the templates

```sh
$ cd $GOPATH/src/arch/../[YOUR OS ARCHITECTURE]/
$ arch [path_to_input_json] [path_to_output_hcl] [arch_name]
```

## Install Linux

Here's how it could look for 64 bits Linux, if you wanted `arch` available globally:

```bash
cd $GOPATH/src/arch/../linux_amd64/ && \
sudo cp arch /usr/local/bin && \
sudo chmod 755 /usr/local/bin/arch && arch -version
```

## Install OSX

Here's how it could look for 64 bits Darwin, if you wanted `arch` available globally:

```bash
cd $GOPATH/src/arch/../darwin_amd64/ && \
sudo cp arch /usr/local/bin && \
sudo chmod 755 /usr/local/bin/arch && arch -version
```
