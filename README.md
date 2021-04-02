# image-admin

## registry

Generate password

``` bash
htpasswd -Bbn tosone testpassword > htpasswd
mkdir auth && mv htpasswd auth
```

Start docker registry

``` bash
docker run -d \
  -p 5000:5000 \
  --restart=always \
  --name registry \
  -v "$(pwd)"/auth:/auth \
  -e "REGISTRY_AUTH=htpasswd" \
  -e "REGISTRY_AUTH_HTPASSWD_REALM=Registry Realm" \
  -e REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd \
  registry:2
```

## notary

``` bash
git clone https://github.com/theupdateframework/notary.git

# install cfssl tool from here https://github.com/cloudflare/cfssl
# cfssl and cfssljson is needed

cd notary/fixtures && ./regenerateTestingCerts.sh

docker-compose up
```

## setting content trust

``` bash
export DOCKER_CONTENT_TRUST_SERVER=https://localhost:4443

docker trust key generate sample_signer
```

## sign image

``` bash
docker trust signer add --key sample_signer.pub sample_signer localhost:5000/release/alpine:3.12
docker trust sign localhost:5000/release/alpine:3.12
```

People can verify image with the sample_signer.pub from the notary server.
