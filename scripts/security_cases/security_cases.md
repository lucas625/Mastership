# Security Cases

This folder contains the scripts for building each of the securty cases.

## Table of Contents

- [Security Cases](#security-cases)
  - [Table of Contents](#table-of-contents)
  - [Cases](#cases)
    - [All not meshed](#all-not-meshed)
    - [Running for Istio](#running-for-istio)
    - [Running for Linkerd](#running-for-linkerd)

## Cases

### All not meshed

```bash
./case_all.sh no-mesh
```

### Running for Istio

- Case all meshed

```bash
./case_all.sh istio
```

- Case Security Checker/Experimenter meshed & Calculator not meshed

```bash
./case_internals_meshed_calculator_not_meshed.sh istio
```

- Case Security Checker/Experimenter not meshed & Calculator meshed

```bash
./case_internals_not_meshed_calculator_meshed.sh istio
```

- Case Mastership not meshed & External Security Checker meshed

```bash
./case_mastership_not_meshed_external_meshed.sh istio
```

- Case Mastership meshed External Security Checker not meshed

```bash
./case_mastership_meshed_external_not_meshed.sh istio
```

### Running for Linkerd

- Case all meshed

```bash
./case_all.sh linkerd
```

- Case Security Checker/Experimenter meshed & Calculator not meshed

```bash
./case_internals_meshed_calculator_not_meshed.sh linkerd
```

- Case Security Checker/Experimenter not meshed & Calculator meshed

```bash
./case_internals_not_meshed_calculator_meshed.sh linkerd
```

- Case Mastership not meshed & External Security Checker meshed

```bash
./case_mastership_not_meshed_external_meshed.sh linkerd
```

- Case Mastership meshed External Security Checker not meshed

```bash
./case_mastership_meshed_external_not_meshed.sh linkerd
```
