# KubeFS

## stage
- [ ] build
- [ ] test 
- [ ] get current version tag
- [ ] increment version 
- [ ] build container
- [ ] push container 


## Prod 
- copy version from stage to prod

# KubeFS Web

## stage
- test 
- build containeri 

## prod
- change urls (sed)
```
You can use sed to do this

For example to replace 'foo' with 'bar' - sed -i 's/foo/bar/g' input_file

- name: Replace credentials
  run: |
    find . -name "*" -exec sed -i "s/__VERSION__/$(cat VERSION.md)/g" {} +
    sed -i 's/__DB-PASSWORD__/${{ secrets.DB_PASSWORD }}/g' db_connection.php
```
- rebuild 
- copy version from stage to prod 


# Increment Version
```bash
a=$(cat env/overlays/stage/version.yml | egrep "image: docker.io/cmwylie19" | sed 's/image: docker.io\/cmwylie19\/kubefs://g' | awk '{$1=$1};1') 

MINOR=$(echo ${a##*.})

MINOR=$(expr $MINOR + 1)


# Patch Version
IMAGE=$(cat env/overlays/stage/version.yml | egrep image: | sed 's/image: //g' | awk '{$1=$1};1')

# Insert backslashes infront of forward slash
NEW_IMAGE=$(echo ${IMAGE%?}$MINOR)

# Insert backslash
NEW_IMAGE=$(echo "$NEW_IMAGE" | sed  's#/#\\/#g')
IMAGE=$(echo "$IMAGE" | sed  's#/#\\/#g')

sed "s/$IMAGE/$NEW_IMAGE/g" env/overlays/stage/version.yml 
```

# Push back to git