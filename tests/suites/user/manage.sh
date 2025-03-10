# Granting and revoking read/write/admins rights for the users.
run_user_grant_revoke() {
	# Echo out to ensure nice output to the test suite.
	echo

	# The following ensures that a bootstrap juju exists.
	file="${TEST_DIR}/test-user-grant-revoke.log"
	ensure "user-grant-revoke" "${file}"

	echo "Check that current user is admin"
	juju whoami --format=json | jq -r '."user"' | check "admin"

	echo "Add user with read rights"
	juju show-user readuser 2>/dev/null || juju add-user readuser
	juju grant readuser read "user-grant-revoke"

	echo "Add user with write rights"
	juju show-user writeuser 2>/dev/null || juju add-user writeuser
	juju grant writeuser write "user-grant-revoke"

	echo "Add user with admin rights"
	juju show-user adminuser 2>/dev/null || juju add-user adminuser
	juju grant adminuser admin "user-grant-revoke"

	echo "Check rights for added users"
	juju show-model "user-grant-revoke" --format=json | jq -r '."user-grant-revoke"."users"."readuser"."access"' | check "read"
	juju show-model "user-grant-revoke" --format=json | jq -r '."user-grant-revoke"."users"."writeuser"."access"' | check "write"
	juju show-model "user-grant-revoke" --format=json | jq -r '."user-grant-revoke"."users"."adminuser"."access"' | check "admin"

	echo "Revoke rights"
	juju revoke readuser read "user-grant-revoke"
	juju revoke writeuser write "user-grant-revoke"
	juju revoke adminuser admin "user-grant-revoke"

	echo "Check rights for added users after revoke"
	juju show-model "user-grant-revoke" --format=json | jq -r '."user-grant-revoke"."users"."readuser"."access"' | check null
	juju show-model "user-grant-revoke" --format=json | jq -r '."user-grant-revoke"."users"."writeuser"."access"' | check "read"
	juju show-model "user-grant-revoke" --format=json | jq -r '."user-grant-revoke"."users"."adminuser"."access"' | check "write"

	destroy_model "user-grant-revoke"
}

# Granting and revoking read/write/admins rights for external users including
# the special everyone@external user.
run_user_grant_revoke_external() {
	# Echo out to ensure nice output to the test suite.
	echo

	# The following ensures that a bootstrap juju exists.
	file="${TEST_DIR}/test-user-grant-revoke-external.log"
	ensure "user-grant-revoke-external" "${file}"

	echo "Check that current user is admin"
	juju whoami --format=json | jq -r '."user"' | check "admin"

	echo "Check that the everyone@external user has been created on bootstrap"
	juju show-user everyone@external --format=json | jq -r '."user-name"' | check "everyone@external"
	echo "Check that the everyone@external user has no permissions"
	juju users --format=json | jq -r '.[] | select(."user-name"=="everyone@external") | ."access"' | check ""

	echo "Add an external user with no permission"
	# Grant the user login permissions and immediately revoke them. Granting an
	# external user permissions is the only way to add the user.
	juju grant nopermuser@external login
	juju revoke nopermuser@external login
	echo "Add an external user with login permissions"
	juju grant loginuser@external login
	echo "Add an external user with superuser permissions"
	juju grant superuser@external superuser

	echo "Check rights for added users"
	juju users --format=json | jq -r '.[] | select(."user-name"=="nopermuser@external") | ."access"' | check ""
	juju users --format=json | jq -r '.[] | select(."user-name"=="loginuser@external") | ."access"' | check "login"
	juju users --format=json | jq -r '.[] | select(."user-name"=="superuser@external") | ."access"' | check "superuser"

	echo "Grant everyone@external login permissions"
	juju grant everyone@external "login"

	echo "Check that external users inherit the permissions"
	juju users --format=json | jq -r '.[] | select(."user-name"=="everyone@external") | ."access"' | check "login"
	juju users --format=json | jq -r '.[] | select(."user-name"=="nopermuser@external") | ."access"' | check "login"
	juju users --format=json | jq -r '.[] | select(."user-name"=="loginuser@external") | ."access"' | check "login"
	juju users --format=json | jq -r '.[] | select(."user-name"=="superuser@external") | ."access"' | check "superuser"

	echo "Revoke login permission of everyone@external"
	juju revoke everyone@external login

	echo "Check that external users have their original permissions"
	juju users --format=json | jq -r '.[] | select(."user-name"=="everyone@external") | ."access"' | check ""
	juju users --format=json | jq -r '.[] | select(."user-name"=="nopermuser@external") | ."access"' | check ""
	juju users --format=json | jq -r '.[] | select(."user-name"=="loginuser@external") | ."access"' | check "login"
	juju users --format=json | jq -r '.[] | select(."user-name"=="superuser@external") | ."access"' | check "superuser"

	echo "Remove added users"
	juju remove-user -y nopermuser@external
	juju remove-user -y loginuser@external
	juju remove-user -y superuser@external

	destroy_model "user-grant-revoke-external"
}

# Disabling and enabling users.
run_user_disable_enable() {
	# Echo out to ensure nice output to the test suite.
	echo

	# The following ensures that a bootstrap juju exists.
	file="${TEST_DIR}/test-user-disable-enable.log"
	ensure "user-disable-enable" "${file}"

	echo "Check that current user is admin"
	juju whoami --format=json | jq -r '."user"' | check "admin"

	echo "Add testuser"
	juju show-user testuser 2>/dev/null || juju add-user testuser
	juju grant testuser read "user-disable-enable"

	echo "Disable testuser"
	juju disable-user testuser

	echo "Check testuser is disabled"
	juju show-user testuser --format=json | jq -r '."disabled"' | check true

	echo "Enable testuser"
	juju enable-user testuser

	echo "Check testuser is enabled"
	juju show-user testuser --format=json | jq -r '."disabled"' | check null

	destroy_model "user-disable-enable"
}

# Granting and revoking login/add-model/superuser rights for the controller access.
run_user_controller_access() {
	# Echo out to ensure nice output to the test suite.
	echo

	# The following ensures that a bootstrap juju exists.
	file="${TEST_DIR}/test-user-controller-access.log"
	ensure "user-controller-access" "${file}"

	echo "Check that current user is admin"
	juju whoami --format=json | jq -r '."user"' | check "admin"

	echo "Add user with login rights"
	juju show-user junioradmin 2>/dev/null || juju add-user junioradmin

	echo "Add user with superuser rights"
	juju show-user senioradmin 2>/dev/null || juju add-user senioradmin
	juju grant senioradmin superuser

	echo "Check rights for added users"
	juju users --format=json | jq -r '.[] | select(."user-name"=="junioradmin") | ."access"' | check "login"
	juju users --format=json | jq -r '.[] | select(."user-name"=="senioradmin") | ."access"' | check "superuser"

	echo "Revoke rights"
	juju revoke junioradmin login
	juju revoke senioradmin superuser

	echo "Check rights for added users after revoke"
	juju users --format=json | jq -r '.[] | select(."user-name"=="junioradmin") | ."access"' | check ""
	juju users --format=json | jq -r '.[] | select(."user-name"=="senioradmin") | ."access"' | check "login"

	destroy_model "user-controller-access"
}

# Removing users.
run_user_remove() {
	# Echo out to ensure nice output to the test suite.
	echo

	# The following ensures that a bootstrap juju exists.
	file="${TEST_DIR}/test-user-remove.log"
	ensure "user-remove" "${file}"

	echo "Check that current user is admin"
	juju whoami --format=json | jq -r '."user"' | check "admin"

	echo "Add testuser2"
	juju show-user testuser2 2>/dev/null || juju add-user testuser2

	users=$(juju users)
	check_contains "${users}" testuser2

	echo "Remove testuser2"
	juju remove-user -y testuser2

	users=$(juju users)
	check_not_contains "${users}" testuser2

	destroy_model "user-remove"
}

test_user_manage() {
	if [ -n "$(skip 'test_user_manage')" ]; then
		echo "==> SKIP: Asked to skip user manage tests"
		return
	fi

	(
		set_verbosity

		cd .. || exit

		run "run_user_grant_revoke"
		run "run_user_grant_revoke_external"
		run "run_user_disable_enable"
		run "run_user_controller_access"
		run "run_user_remove"
	)
}
