START TRANSACTION;
INSERT into userpassword
(hashedpassword, lastchanged, active, userid)
VALUES
    (MD5("password1"), sysdate(), true, 1)
;

INSERT into userpassword
(hashedpassword, lastchanged, active, userid)
VALUES
    (MD5("password2"), sysdate(), true, 1)
;

UPDATE userpassword
SET
    lastchanged = sysdate(),
    active = false
WHERE
        userid = 1
    and active = true
;

INSERT into userpassword
(hashedpassword, lastchanged, active, userid)
VALUES
    (MD5("password3"), sysdate(), true, 1)
;

COMMIT;