# Generated by Django 4.0.5 on 2022-06-19 08:13

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('ngouzo_api', '0002_parentprofile_and_more'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='user',
            name='first_name',
        ),
        migrations.RemoveField(
            model_name='user',
            name='is_active',
        ),
        migrations.RemoveField(
            model_name='user',
            name='last_name',
        ),
        migrations.AddField(
            model_name='user',
            name='last_login',
            field=models.DateTimeField(blank=True, null=True, verbose_name='last login'),
        ),
    ]
