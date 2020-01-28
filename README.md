# gokit-microservices
gokit microservices sample code


`curl -XPOST -d'{"str":"Shivam"}' localhost:8080/count`

// output

{"length":6}

`curl -XPOST -d'{"str":"Shivam"}' localhost:8080/uppercase`

// output

{"str":"SHIVAM"}

# open below URL in browser
http://localhost:8080/metrics

# for prometheus
install binary and run application after configuring prometheus.yml
- add new job with you server ips

    - job_name: "microservice"

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.

    static_configs:
      - targets: ["localhost:8080"] // golang server ip

# install grafana for better visualization of prometheus
