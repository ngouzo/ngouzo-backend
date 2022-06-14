# from xml.etree.ElementInclude import include
from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    path('ngouzo/', include("ngouzo_api.urls")),
    path('admin/', admin.site.urls),
]
