CREATE TABLE recorder(
                         recorder_ID bigserial not null ,
                         device_ID int not null,
                         N int,
                         Mqtt varchar(255),
                         Invid varchar(100),
                         Msg_id varchar(100),
                         Text varchar(100),
                         Context varchar(100),
                         Class varchar(100),
                         Level integer,
                         Area varchar(100),
                         Block  varchar(100),
                         Type_ varchar(100),
                         Bit_ varchar(100),
                         Invert_bi varchar(100)
);

CREATE TABLE device(
                       device_ID bigserial not null,
                       unig_guid varchar(255) unique,
                       CONSTRAINT device_PK PRIMARY KEY (device_ID)
);

CREATE TABLE file_name(
                          name varchar(255) not null,
                          flag_error varchar(10) not null, /*сообщение об успешности прочтения файла*/
                          file_ID bigserial not null,
                          CONSTRAINT file_PK PRIMARY KEY (file_ID)
);

CREATE TABLE file_and_device(
                                device_ID int,
                                file_and_device_ID bigserial not null,
                                file_ID int not null
);

ALTER TABLE recorder ADD CONSTRAINT FK_recorder
    FOREIGN KEY (device_ID)
        REFERENCES device(device_ID)
;

ALTER TABLE file_and_device ADD CONSTRAINT FK_file_and_device
    FOREIGN KEY (device_ID)
        REFERENCES device(device_ID)
;

ALTER TABLE file_and_device ADD CONSTRAINT FK_file
    FOREIGN KEY (file_ID)
        REFERENCES file_name(file_ID)
;