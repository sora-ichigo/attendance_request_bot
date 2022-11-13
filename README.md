友人への講義の出席依頼を Line Bot で自動化する

## Design

```sh
cmd
└── request_syspro_attendance_job # シスプロの出席を依頼するJob
    ├── main.go
    └── run.go
```

これらのエンドポイントを CloudRun へデプロイし、CloudScheduler で毎週の講義前に定期実行する。
