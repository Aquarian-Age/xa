#
# spec file for package ccal-all
#
# Copyright (c) 2021 SUSE LLC
#
# All modifications and additions to the file contributed by third parties
# remain the property of their copyright owners, unless otherwise agreed
# upon. The license for this file, and modifications and additions to the
# file, is the same license as for the pristine package itself (unless the
# license for the pristine package is not an Open Source License, in which
# case the license is the MIT License). An "Open Source License" is a
# license that conforms to the Open Source Definition (Version 1.9)
# published by the Open Source Initiative.

# Please submit bugfixes or comments via https://bugs.opensuse.org/
#


Name:           ccal-all
Version:        0.6.9
Release:        6
Summary:        ccal-all
License:        Me
URL:            https://github.com/Aquarian-Age/ccal/releases 
Source0:       ccal-all.tar.gz 
BuildRoot:      %{_tmppath}-%{version}-%{release}-build
BuildRequires: gtk2-devel >= 2.0.0
%description
Chinese Lunar Calendar QiMen XieJiBianFangShu XiaoLiuRen


%prep
rm -rf %{name}
rm -rf %{buildroot}

%build
tar xzvf ../SOURCES/ccal-all.tar.gz -C %{buildroot}/
cd %{buildroot}
go build -o ccal-all -ldflags="-s -w" -tags timetzdata

%install
mkdir -p %{buildroot}/usr/local/bin/
cp -f %{buildroot}/ccal-all %{buildroot}/usr/local/bin/ccal-all
mkdir -p %{buildroot}/usr/local/share/applications/
mkdir -p %{buildroot}/usr/local/share/icons/
cp -f %{buildroot}/fire.svg %{buildroot}/usr/local/share/icons/fire.svg
cp -f %{buildroot}/ccal-all.desktop %{buildroot}/usr/local/share/applications/ccal-all.desktop

rm -rf %{buildroot}/*.go
rm -rf %{buildroot}/go.*
rm -rf %{buildroot}/ccal-*
rm -rf %{buildroot}/*.svg
rm -rf %{buildroot}/cal.tis

%files
%defattr(-,root,root)
/usr/local/bin/ccal-all
/usr/local/share/applications/ccal-all.desktop
/usr/local/share/icons/fire.svg

%changelog
qiMen yueJiang
