# Running Migration for Gorm using Atlas

This guide will walk you through the steps to run migration for Gorm using Atlas.

## Prerequisites

- Golang installed on your machine
- Atlas installed and set up.

```bash
curl -sSf https://atlasgo.sh | sh
```

## Steps

1. **Set Up Gorm Models**

    - Create your models in go.
    - Open `src/migration/main.go`, import your models there.

2. **Run Migration**

    - Run diff to generate the migration
      ```
         atlas migrate diff --env local
      ```
    - Apply the migration

   ```
        atlas migrate apply --env local
   ```

   Or using `make` (default local env)

   ```bash
   # generate diff
   make diff

   # Apply
   make apply
   ```

3. **Verify Migration**

    - After running the migration, verify that the database schema has been updated as expected.

4. **Handle Errors**

    - If there are any errors during the migration process, troubleshoot and fix them before proceeding.

5. **Commit and Push Changes**
    - Once the migration is successful, commit your migration files and push them to your version control system.

## Additional Considerations

- Test your migrations in a development environment before applying them to production.
- Follow best practices for version controlling your migration files.

Atlas documentation: https://atlasgo.io/versioned/intro

That's it! You have now successfully run migration for Gorm using Atlas.

# Gen seed data

1. Create template file in `seed-tmpl/` with the file name extension ".tmpl.sql" following the go template syntax [https://pkg.go.dev/text/template](https://pkg.go.dev/text/template).
2. Create seed data in `seed-data/` with the file name extension ".json".
3. Run `make seed tpl=./seed-tmpl/[filename].tmpl.sql seed=./seed-data/[data].json` to generate seed data.
4. Run `make apply-seed` to apply seed data to database.

## Troubleshooting

1. If you get an error like this:

   ```bash
   Error: Error 1146: Table 'atlas_test.users' doesn't exist
   ```

   It means that the table you are trying to seed does not exist in the database. You can fix this by running the migration first.

2. If you adjust the data in the seed file, you may need to delete the existing seed data in the database before applying the seed data again. Ensure that the data in the seed file is consistent with the data in the database.

3. The timestamp of the seed data is based on the time when the seed data is generated. And must beyond the timestamp of the last migration. If you want to apply the seed data to the database, you need to make sure that the timestamp of the seed data is greater than the timestamp of the last migration. You can check the timestamp of the last migration in the `atlas_migrations` table. If not you can delete the seed data in the database and regenerate the seed data.
