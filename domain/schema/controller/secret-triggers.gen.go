// Code generated by triggergen. DO NOT EDIT.

package controller

import (
	"fmt"

	"github.com/juju/juju/core/database/schema"
)


// ChangeLogTriggersForSecretBackendRotation generates the triggers for the 
// secret_backend_rotation table.
func ChangeLogTriggersForSecretBackendRotation(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for SecretBackendRotation
CREATE TRIGGER trg_log_secret_backend_rotation_insert
AFTER INSERT ON secret_backend_rotation FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for SecretBackendRotation
CREATE TRIGGER trg_log_secret_backend_rotation_update
AFTER UPDATE ON secret_backend_rotation FOR EACH ROW
WHEN 
	NEW.next_rotation_time != OLD.next_rotation_time 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for SecretBackendRotation
CREATE TRIGGER trg_log_secret_backend_rotation_delete
AFTER DELETE ON secret_backend_rotation FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

