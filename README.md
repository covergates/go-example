# go-example

This is a example repository to show how to generate and upload coverage report to [Covergates](https://covergates.com/) for `Go`.

To get started, please download the prebuilt binary from [release page](https://github.com/covergates/covergates/releases).
Unzip the binaries to your `/path/to/bin`.

## Fork the Repository

Please fork the repository at first by click the **Fork** on the top-right of the repository page.

![fork](https://raw.githubusercontent.com/covergates/go-example/master/assets/fork.png)

After that, clone the repository to your workspace with:

```
git clone https://github.com/<your account>/go-example.git
cd go-example
```

## Generate Coverage Report

It is assume that you have installed `Go`. If you haven't installed it, refer to [installation guide](https://golang.org/doc/install).

Run `go test`:

```
go test -coverprofile=coverage.out .
```

You will see below standard outputs:

```
ok      go-example      0.001s  coverage: 50.0% of statements
```

The coverage report is also generated, which is `coverage.out` with below content:

```
mode: set
go-example/main.go:5.24,7.2 1 1
go-example/main.go:9.13,11.2 1 0
```

## Activate Repository

Before uploading report, you need to activate repository on **Covergates**.
Visit [https://covergates.com/repo](https://covergates.com/repo) and click **ACTIVATE** right to `go-example` repository.

![repository](https://raw.githubusercontent.com/covergates/go-example/master/assets/repository.png)

> :information_desk_person: You require login **Covergate** with your GitHub Account to continue

Visit repository setting page [https://covergates.com/report/github/your-account/go-example/setting](https://covergates.com/report/github/your-account/go-example/setting)

You could find the **Report ID** which is used to upload report as below image shown:

![setting](https://github.com/covergates/go-example/blob/master/assets/setting.png?raw=true)

## Upload Report

Go back to your `go-example` workspace, run:

```
export REPORT_ID="your Report ID"
covergates upload --type go ./coverage.out
```

You should see below message is uploading success.

```
2020/08/09 22:44:47 upload commit 5b9f6246e7afe2d86c52d5673e90a1d6e172022c, go
2020/08/09 22:44:48 ok
```

Visit [https://covergates.com/report/github/your-account/go-example](https://covergates.com/report/github/your-account/go-example) to see the report.
