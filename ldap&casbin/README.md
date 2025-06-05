![alt text](image.png)
![alt text](image-1.png)

docker run -p 389:389 --name openldap --env LDAP_ORGANISATION="Example Org" --env LDAP_DOMAIN="example.org" --env LDAP_ADMIN_PASSWORD=admin -d osixia/openldap
