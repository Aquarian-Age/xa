#
# spec file for package groff-utf8-1.22.tar.gz
#
# Copyright (c) 2018 SUSE LINUX GmbH, Nuernberg, Germany.
#
# All modifications and additions to the file contributed by third parties
# remain the property of their copyright owners, unless otherwise agreed
# upon. The license for this file, and modifications and additions to the
# file, is the same license as for the pristine package itself (unless the
# license for the pristine package is not an Open Source License, in which
# case the license is the MIT License). An "Open Source License" is a
# license that conforms to the Open Source Definition (Version 1.9)
# published by the Open Source Initiative.

# Please submit bugfixes or comments via http://bugs.opensuse.org/
#


Name:          ccal-auth
Version:       0.6.9 
Release:       0
Summary:       Chinese Lunar Calendar
License:       me
Group:         Application
Url:           https://github.com/Aquarian-Age/ccal/releases 
Source0:       ccal-auth.tar.gz 

BuildRoot:      %{_tmppath}/%{name}-%{version}-build
BuildRequires: glibc >= 2.3.2
BuildRequires: gtk2-devel >= 2.0.0

%description
Chinese Lunar Calendar


%prep

%build
tar xzvf ../SOURCES/ccal-auth.tar.gz -C ${RPM_BUILD_ROOT}/
cd ${RPM_BUILD_ROOT}/
go build -o ccal-auth -ldflags="-s -w" -tags timetzdata

%install
mkdir -p ${RPM_BUILD_ROOT}/usr/local/bin/
cp -f ${RPM_BUILD_ROOT}/ccal-auth ${RPM_BUILD_ROOT}/usr/local/bin/ccal-auth
mkdir -p ${RPM_BUILD_ROOT}/usr/local/share/applications/
mkdir -p ${RPM_BUILD_ROOT}/usr/local/share/icons/
cp -f ${RPM_BUILD_ROOT}/aquarius.svg ${RPM_BUILD_ROOT}/usr/local/share/icons/aquarius.svg
cp -f ${RPM_BUILD_ROOT}/ccal-auth.desktop ${RPM_BUILD_ROOT}/usr/local/share/applications/ccal-auth.desktop

rm -rf ${RPM_BUILD_ROOT}/*.go
rm -rf ${RPM_BUILD_ROOT}/go.*
rm -rf ${RPM_BUILD_ROOT}/ccal-*
rm -rf ${RPM_BUILD_ROOT}/aquarius.svg

%files
%defattr(-,root,root)
/usr/local/bin/ccal-auth
/usr/local/share/applications/ccal-auth.desktop
/usr/local/share/icons/aquarius.svg


%changelog
