# Generated by Django 3.0.8 on 2020-07-22 05:43

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('todo', '0001_initial'),
    ]

    operations = [
        migrations.AddField(
            model_name='todomodel',
            name='date_created',
            field=models.DateTimeField(auto_now=True),
        ),
    ]
