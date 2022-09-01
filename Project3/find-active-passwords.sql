SELECT
    userpassword.userid,
    userpassword.hashedpassword
FROM
    userpassword
WHERE
    userpassword.active = true
;