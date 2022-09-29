CREATE TABLE gitlab_projects
(
  id  integer  not null unique,
  name    varchar(255) not null,
  description varchar(255) not null,
  path_with_namespace varchar(255) not null,
  web_url varchar(255) not null,
  archived bool not null,
  username varchar(255) not null,
  owner_name varchar(255) not null
);

CREATE TABLE appsec_users
(
  user_id  int  not null unique,
  user_name varchar(255) not null,
  user_phone varchar(255) not null,
  user_email varchar(255) not null
);

CREATE TABLE user_projects
(
  user_id  int  not null,
  git_url varchar(255) not null
);

BEGIN;	
ALTER TABLE gitlab_projects ADD namespace_name varchar(255);
ALTER TABLE gitlab_projects ADD namespace_id integer;	
COMMIT;

CREATE TABLE appsechub_defects
(
  id  integer  not null unique,
	appid    integer not null,
  summary    varchar(255) not null,
	priority_id varchar(255) not null,
	jira_link varchar(255) not null,
	status varchar(255) not null,
	created varchar(255)
);

CREATE TABLE appsechub_projects
(
  id  integer  not null unique,
	name    varchar(255) not null,
  code    varchar(255) not null,
	project_link varchar(255) not null,
	defects_total integer  not null,
	issues_total integer  not null,
	ssdl_total integer  not null,
	codebase_size integer  not null
);

CREATE TABLE appsechub_artifacts
(
  id  integer  not null unique,
	repository_url varchar(255) not null,
	artifact    varchar(255) not null
);

CREATE TABLE appsechub_codebase
(
  id  integer  not null unique,
	name varchar(255) not null,
	link    varchar(255) not null,
  branch    varchar(255) not null,
  active    bool not null,
  appId    integer not null
);


CREATE TABLE appfarm_services
(
  id  varchar(512)  not null unique,
  environment varchar(512) not null,
  information_system_id    varchar(512) not null,
  name   varchar(512) not null,
  description text not null,
  project_path varchar(512) not null,
  service_type varchar(512) not null,
  owner varchar(512) not null,
  portal_link varchar(512) not null,
  update timestamp not null
);

CREATE TABLE jira_report_v2
(
  id  integer  not null unique,
  key varchar(255) not null,
  name    varchar(255) not null,
  summary varchar(255) not null,
  status    varchar(255) not null,
  labels     jsonb not null default '{}'::jsonb,
  created    timestamp not null,
  updated    timestamp not null,
  time_taken integer
);

ALTER TABLE jira_report ADD summary varchar(255);	

CREATE TABLE jira_report_v3
(
  id  integer  not null unique,
  key varchar(255) not null,
  name    varchar(255) not null,
  summary varchar(255) not null,
  status    varchar(255) not null,
  labels     jsonb not null default '{}'::jsonb,
  components varchar(255),
  created    timestamp not null,
  updated    timestamp not null,
  time_taken integer
);

ALTER TABLE jira_report_v2 ADD resolutiondate timestamp;	
ALTER TABLE jira_report_v3 ADD resolutiondate timestamp;	

CREATE TABLE jira_work_time
(
  id  integer  not null unique,
  issue_id integer,
  name    varchar(255),
  comment varchar(1024),
  created    timestamp,
  updated    timestamp,
  started timestamp,
  time_spent varchar(255),
  time_spent_seconds integer
);

ALTER TABLE jira_report_v2 ADD url varchar(255);
ALTER TABLE jira_report_v3 ADD url varchar(255);


BEGIN;
ALTER TABLE appfarm_services
    ALTER COLUMN description TYPE text,
    ALTER COLUMN description SET NOT NULL

 ALTER TABLE gitlab_projects
    ALTER COLUMN description TYPE text,
    ALTER COLUMN description SET NOT NULL
 COMMIT;      

 ALTER TABLE appfarm_services ADD internal_url varchar(255);
 ALTER TABLE appfarm_services ADD public_url varchar(255);
