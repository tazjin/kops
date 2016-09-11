# Delegating a subdomain to Amazon Route53
## Subdomain creation without migrating the parent domain

You can create a subdomain that uses Amazon Route 53 as the DNS service without migrating the parent domain from another DNS service.

Example:

Your main domain is `example.com` but you want to create a subdomain of `k8s.example.com`.

### Check current status of your domain:

> dig ns example.com

``` 
;; QUESTION SECTION:
;example.com.			IN	NS
 
;; ANSWER SECTION:
example.com.		3600	IN	NS	ns3.somensserver.com.
```

### Create subdomain

You want to keep those records, now lets create the subdomain.

On your `Route 53` create the subdomain:

`Create hosted zone`

Fill up the box `Domain Name:` with your subdomain : k8s.example.com

`Route 53` should generate your NS server like:

```
;; ANSWER SECTION:
ns-613.awsdns-13.net.
ns-75.awsdns-04.com.
ns-1022.awsdns-35.co.uk.
ns-1149.awsdns-27.org.
```

With your registrar add those NS server to your subdomain.

The result should be.

>dig ns k8s.example.com

```
;; ANSWER SECTION:
k8s.example.com.		172800	IN	NS	ns-613.awsdns-13.net.
k8s.example.com.		172800	IN	NS	ns-75.awsdns-04.org.
k8s.example.com.		172800	IN	NS	ns-1022.awsdns-35.com.
k8s.example.com.		172800	IN	NS	ns-1149.awsdns-27.co.uk.
```

Wait until the DNS records have propagated.
