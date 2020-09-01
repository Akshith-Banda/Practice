from django.db import models
from django.utils import timezone
from django.contrib.auth.models import User

class TodoModel(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE, blank=True, null=True)
    text = models.CharField(max_length=100)
    completed = models.BooleanField(default=False)
    date_created = models.DateTimeField(default=timezone.now())

    def __str__(self):
        return self.text
