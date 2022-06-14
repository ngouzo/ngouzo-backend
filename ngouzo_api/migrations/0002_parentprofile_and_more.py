# Generated by Django 4.0.5 on 2022-06-14 12:23

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('ngouzo_api', '0001_initial'),
    ]

    operations = [
        migrations.CreateModel(
            name='ParentProfile',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('picture', models.ImageField(blank=True, upload_to='uploads/')),
                ('sex', models.CharField(blank=True, default='', max_length=10)),
                ('birth_date', models.DateField(blank=True)),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('updated', models.DateTimeField(auto_now_add=True)),
            ],
        ),
        migrations.RenameField(
            model_name='employeefinancial',
            old_name='teach_id',
            new_name='employee_id',
        ),
        migrations.RenameField(
            model_name='employeepaymenthistory',
            old_name='teach_id',
            new_name='employee_id',
        ),
        migrations.RemoveField(
            model_name='userprofile',
            name='is_parent',
        ),
        migrations.AddField(
            model_name='user',
            name='is_parent',
            field=models.BooleanField(blank=True, default=False),
        ),
        migrations.CreateModel(
            name='TeacherPost',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('post', models.TextField(blank=True, default='')),
                ('sent_at', models.DateTimeField(auto_now_add=True)),
                ('removed_at', models.DateTimeField(auto_now_add=True)),
                ('updated', models.DateTimeField(auto_now_add=True)),
                ('parent_id', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='ngouzo_api.userprofile')),
                ('teacher_id', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='ngouzo_api.user')),
            ],
        ),
        migrations.CreateModel(
            name='PostCommunication',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('post', models.TextField(blank=True, default='')),
                ('sent_at', models.DateTimeField(auto_now_add=True)),
                ('removed_at', models.DateTimeField(auto_now_add=True)),
                ('updated', models.DateTimeField(auto_now_add=True)),
                ('school_id', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='ngouzo_api.user')),
            ],
        ),
        migrations.CreateModel(
            name='ParentSchoolStudent',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('updated', models.DateTimeField(auto_now_add=True)),
                ('parent_id', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='ngouzo_api.user')),
                ('student_id', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='ngouzo_api.students')),
            ],
        ),
        migrations.CreateModel(
            name='ParentSchool',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('updated', models.DateTimeField(auto_now_add=True)),
                ('parent_profile_id', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='ngouzo_api.parentprofile')),
                ('school_id', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='ngouzo_api.user')),
            ],
        ),
        migrations.AddField(
            model_name='parentprofile',
            name='parent_id',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='ngouzo_api.user'),
        ),
    ]