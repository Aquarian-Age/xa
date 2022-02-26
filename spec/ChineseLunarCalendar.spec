#
# spec file for package ChineseLunarCalendar
#
# Copyright (c) 2022 SUSE LLC
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


Name:           ChineseLunarCalendar
Version:        1.0.0
Release:        0
Summary:        Constellation ChineseLunarCalendar
License:        MIT
URL:            https://github.com/Aquarian-Age/xa
Source0:        ccal.tar.gz

%if 0%{?suse_version} >= 1150
BuildRequires: golang(API) = 1.15
%else
BuildRequires: golang(API) = 1.17
%endif

BuildRequires:  gcc 
BuildRequires:  pkg-config 
BuildRoot:      %{_tmppath}/%{name}-%{version}-build   

%description
ChineseLunarCalendar

%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/ccal.tar.gz -C %{_builddir}/%{name}-%{version}

%build
cd %{_builddir}/%{name}-%{version}

go build -o ChineseLunarCalendar -mod=vendor -tags tempdll -ldflags="-s -w -X 'main.Version=1.0.0t' -X 'main.GoVersion=Go Version: go1.17.6' -X 'main.Mail=bGlhbmd6aTIwMjFAeWFuZGV4LmNvbQo='" -trimpath .

%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/
install -d %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/
install -m0755 %{_builddir}/%{name}-%{version}/ChineseLunarCalendar %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/ChineseLunarCalendar

mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/applications/
install -d %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}
install -m0755 %{_builddir}/%{name}-%{version}/ChineseLunarCalendar.png %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/ChineseLunarCalendar.png
install -m0755 %{_builddir}/%{name}-%{version}/ChineseLunarCalendar.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/applications/ChineseLunarCalendar.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/*.go
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/go.*
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/ChineseLunarCalendar*
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/ChineseLunarCalendar.png
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
%{_bindir}/ChineseLunarCalendar
%{_datadir}/applications/ChineseLunarCalendar.desktop
%{_datadir}/icons/ChineseLunarCalendar.png


%changelog
