CREATE PROCEDURE `createSessionSurvey` ()
BEGIN
create procedure  procedcreateSurveyFromSession()
begin
DECLARE sid INT;
DECLARE done INT DEFAULT 0;
DECLARE rs CURSOR FOR SELECT sessionId FROM SAP.Session;
DECLARE CONTINUE HANDLER FOR SQLSTATE '02000' SET Done = 1;  
OPEN rs;
FETCH NEXT FROM rs INTO sid;
REPEAT
SELECT @RS;
IF NOT done THEN
INSERT INTO sap.survey_info (sessionid) VALUES (sid);
END IF;
FETCH NEXT FROM rs INTO sid;
UNTIL done END REPEAT;
CLOSE rs;
END
